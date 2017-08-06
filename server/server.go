package server

import (
	"database/sql"
	"encoding/binary"
	"io"

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

func (s *Server) Run() error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Errorw("Recovered main server", "details", r)
		}
	}()

	world := actor.Spawn(actor.FromInstance(&entities.World{}))

	listen, err := net.Listen("tcp", s.config.Common.Host)
	if err != nil {
		return err
	}

	s.logger.Infow("Started server", "host", s.config.Common.Host)

	for {
		client, err := listen.Accept()
		if err != nil {
			s.logger.Errorw("Error accept connection", "err", err)
			continue
		}
		defer client.Close()

		var (
			packetSender = actor.Spawn(router.NewRoundRobinPool(5).WithInstance(&entities.PacketSender{
				Client:  client,
				Network: network.NewNetwork(),
				Logger:  s.logger,
			}))
			player = actor.Spawn(actor.FromInstance(&entities.Player{
				World:        world,
				PacketSender: packetSender,
			}))
			packetReader = actor.Spawn(router.NewRoundRobinPool(5).WithInstance(&entities.PacketReader{
				World:        world,
				Player:       player,
				PacketSender: packetSender,
				Logger:       s.logger,
			}))
			connectReader = actor.Spawn(router.NewRoundRobinPool(5).WithInstance(&entities.ConnectReader{
				Client:       client,
				PacketReader: packetReader,
				Network:      network.NewNetwork(),
				Logger:       s.logger,
			}))
		)

		packetSender.Tell(entities.SendPacket{
			Packet: (&out.Date{}).SetCurrentTime(),
		})

		var (
			lenPacket int
			bb        *bytebufferpool.ByteBuffer
		)
		for {
			if lenPacket == 0 {
				bb = bytebufferpool.Get()
			}

			bufTemp := make([]byte, 1536)
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

			if lenPacket == 0 {
				lenPacket = int(binary.BigEndian.Uint16(bb.Bytes()[0:2]))
			}

			if lenPacket < int(ln) {
				continue
			}

			bb.Set(bb.Bytes()[2:])

			connectReader.Tell(entities.ReadPacket{
				Len: lenPacket,
				Buf: bb.Bytes(),
			})

			// Clear things
			bytebufferpool.Put(bb)
			lenPacket = 0
		}
	}
}
