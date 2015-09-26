package land

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/pill"
	log "github.com/Sirupsen/logrus"

	"bytes"
	"net"
)

// ClientLive method for accept new connection from socket
func ClientLive(buffers Buffers, conf core.Config, c net.Conn) {
	defer core.ErrorNetworkHandler(c)

	buffer := bytes.NewBuffer([]byte{})
	player := entitie.NewPlayer(c)
	pillInit := pill.NewPill()

	for getBytes := range buffers.GetReadChannel() {
		buffer.Write(getBytes)

		if conf.Base.Test {
			log.Panic("Client is break =_=")
		} else {
			log.WithField("bytes", buffer.Bytes()).Info("Print message from client")

			if buffer.Len() >= 8 {
				opcodes, err := pillInit.Decrypt(buffer.Bytes(), *player)
				if err != nil {
					log.WithError(err).Panic("Error in pill decrypt")
				}

				if opcodes != nil {
					for _, v := range opcodes {
						pillEncrypt, err := pillInit.Encrypt(pillInit.SetOpcode(v).GetOutcomingCrumb())
						if err != nil {
							log.WithError(err).Error("Error in pill encrypt")
						}

						buffers.GetWriteChannel() <- pillEncrypt
					}
				}
			} else {
				buffers.GetWriteChannel() <- []byte{0x00, 0x02}
			}
		}

		buffer.Reset()
	}
}
