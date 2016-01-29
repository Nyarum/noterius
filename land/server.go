package land

import (
	"bytes"
	"fmt"

	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/packet"
	"github.com/Nyarum/noterius/robot"
	log "github.com/Sirupsen/logrus"

	"net"
)

// Application struct for project and his variables
type Application struct {
	Config   core.Config
	Database core.Database
}

// ClientLive method for accept new connection from socket
func ConnectHandler(buffers *core.Buffers, conf core.Config, c net.Conn) {
	defer core.ErrorNetworkHandler(c)

	var (
		buffer      *bytes.Buffer   = bytes.NewBuffer([]byte{})
		player      *entitie.Player = entitie.NewPlayer(buffers)
		packetAlloc *packet.Packet  = packet.NewPacket(player)
	)

	// Once send first a packet
	packet, err := packetAlloc.Encode(940)
	if err != nil {
		log.WithError(err).Error("Error in pill encrypt")
	}

	buffers.GetWC() <- packet

	for getBytes := range buffers.GetRC() {
		buffer.Reset()
		buffer.Write(getBytes)

		log.WithField("bytes", fmt.Sprintf("% x", buffer.Bytes())).Info("Print message from client")

		// Ping <-> pong
		if buffer.Len() <= 2 {
			buffers.GetWC() <- []byte{0x00, 0x02}
			continue
		}

		err := packetAlloc.Decode(buffer.Bytes())
		if err != nil {
			log.WithError(err).Error("Error in pill decrypt")
			return
		}
	}
}

// Run method for starting server
func (a *Application) Run() (err error) {
	listen, err := net.Listen("tcp", a.Config.Base.IP+":"+a.Config.Base.Port)
	if err != nil {
		return
	}

	// Init robot factories
	robot := robot.NewRobot()
	for _, factory := range robot.Factories {
		go factory.Process(a.Config)
	}

	for {
		client, err := listen.Accept()
		if err != nil {
			log.WithError(err).Error("Error in accept connection")
			continue
		}

		go func(c net.Conn, conf core.Config) {
			var buffers *core.Buffers = core.NewBuffers()

			defer func() {
				buffers.Close()
				core.ErrorNetworkHandler(c)
			}()

			log.WithFields(log.Fields{
				"address": c.RemoteAddr(),
			}).Info("Client is connected")

			go ConnectHandler(buffers, conf, c)
			go buffers.WriteHandler(c)

			buffers.ReadHandler(c, conf)
		}(client, a.Config)
	}
}
