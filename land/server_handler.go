package land

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/packet"
	"github.com/Nyarum/noterius/support"
	log "github.com/Sirupsen/logrus"
)

// ConnectHandler method for accept new connection from socket
func ConnectHandler(buffers *core.Buffers, app Application, c net.Conn) {
	defer core.ErrorNetworkHandler(c)

	var (
		buffer      *bytes.Buffer      = bytes.NewBuffer([]byte{})
		player      *entitie.Player    = entitie.NewPlayer(buffers)
		database    *database.Database = database.NewDatabase(&app.DatabaseInfo)
		packetAlloc *packet.Packet     = packet.NewPacket(player, database)
	)

	// Once send first a packet
	packet, err := packetAlloc.Encode(support.OP_SERVER_CHAPSTR)
	if err != nil {
		log.WithError(err).Error("Error in packet encode")
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

		opcodes, err := packetAlloc.Decode(buffer.Bytes())
		if err != nil {
			log.WithError(err).Error("Error in packet decode")
			return
		}

		if len(opcodes) == 0 {
			continue
		}

		for _, opcode := range opcodes {
			response, err := packetAlloc.Encode(opcode)
			if err != nil {
				log.WithError(err).Error("Error in packet encode")
				break
			}

			buffers.GetWC() <- response
		}

		if player.Error != nil {
			// Before disconnect
			time.Sleep(time.Second)

			log.WithError(player.Error).Error("Client was rejected by God!")
			return
		}
	}
}
