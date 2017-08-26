package server

import (
	"context"
	"database/sql"
	"encoding/binary"
	"io"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/router"
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/entities"
	"github.com/Nyarum/noterius/network"
	"github.com/Nyarum/noterius/network/out"
	"github.com/valyala/bytebufferpool"
	"go.uber.org/zap"

	"net"
)

type Server struct {
	config   core.Config
	database *sql.DB
	logger   *zap.SugaredLogger
}

func NewServer(config core.Config, database *sql.DB, logger *zap.SugaredLogger) *Server {
	return &Server{
		config:   config,
		database: database,
		logger:   logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Errorw("Recovered main server", "details", r)
		}
	}()

	world := actor.Spawn(actor.FromInstance(entities.NewWorld(s.database)))
	defer world.Stop()

	listen, err := net.Listen("tcp", s.config.Common.Host)
	if err != nil {
		return err
	}

	s.logger.Infow("Started server", "host", s.config.Common.Host)

	tick := time.Tick(1 * time.Second / 2)

	// Graceful shutdown and global ticker
	go func() {
		for {
			select {
			case tm := <-tick:
				world.Tell(entities.GlobalTick{
					Now: tm,
				})
			case <-ctx.Done():
				listen.Close()
			}
		}
	}()

	for {
		client, err := listen.Accept()
		if err != nil {
			if ctx.Err() == context.Canceled {
				return nil
			}

			s.logger.Errorw("Error accept connection", "err", err)
			continue
		}

		s.logger.Infow("New connect", "ip", client.RemoteAddr())

		go s.acceptConnect(ctx, client, world)
	}
}

func (s *Server) acceptConnect(ctx context.Context, client net.Conn, world *actor.PID) {
	var (
		packetSender = actor.Spawn(router.NewRoundRobinPool(5).WithInstance(&entities.PacketSender{
			Client:  client,
			Network: network.NewNetwork(),
			Logger:  s.logger,
		}))
		player = actor.Spawn(actor.FromInstance(&entities.Player{
			DB:           s.database,
			World:        world,
			PacketSender: packetSender,
			Logger:       s.logger,
		}))
		packetReader = actor.Spawn(router.NewRoundRobinPool(5).WithInstance(&entities.PacketReader{
			Client:  client,
			Player:  player,
			Network: network.NewNetwork(),
			Logger:  s.logger,
		}))
	)
	defer func() {
		s.logger.Debugw("Shutdown connection actors", "ip", client.RemoteAddr())

		packetReader.Stop()
		packetSender.Stop()
		player.Stop()
	}()

	player.Tell(entities.RecordTime{
		Time: (&out.Date{}).GetCurrentTime(),
	})

	var (
		lenPacket int
		bb        *bytebufferpool.ByteBuffer
	)
	for {
		if ctx.Err() == context.Canceled {
			break
		}

		if lenPacket == 0 {
			bb = bytebufferpool.Get()
		}

		bufTemp := make([]byte, 4096)
		ln, err := client.Read(bufTemp)
		if err != nil {
			if val, ok := err.(net.Error); ok && val.Timeout() {
				s.logger.Errorw("Client is timeout", "error", err)
			}

			if err == io.EOF {
				s.logger.Errorw("Client is disconnected", "error", err)
			}

			break
		}

		bb.Write(bufTemp[:ln])

		// Func to receive many sub packets in an one main packet
		var manyDataFunc func() bool
		manyDataFunc = func() bool {
			if lenPacket == 0 {
				lenPacket = int(binary.BigEndian.Uint16(bb.Bytes()[0:2]))
			}

			if lenPacket < int(ln) {
				return false
			}

			packetReader.Tell(entities.ReadPacket{
				Len: lenPacket,
				Buf: bb.Bytes(),
			})

			bb.Set(bb.Bytes()[lenPacket:])
			lenPacket = 0

			if bb.Len() != 0 {
				return manyDataFunc()
			}

			return false
		}

		if !manyDataFunc() {
			continue
		}

		bytebufferpool.Put(bb)
	}
}
