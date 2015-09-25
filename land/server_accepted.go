package land

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/pill"

	"bytes"
	"log"
	"net"
)

// ClientLive method for accept new connection from socket
func ClientLive(buffers Buffers, conf core.Config, c net.Conn) {
	buffer := bytes.NewBuffer([]byte{})
	player := entitie.NewPlayer(c)
	pillInit := pill.NewPill()

	for getBytes := range buffers.GetReadChannel() {
		buffer.Write(getBytes)

		if conf.Base.Test {
			panic("Client is break =_=")
		} else {
			log.Printf("Print message from client: % x", buffer.Bytes())

			if buffer.Len() >= 8 {
				opcodes, err := pillInit.Decrypt(buffer.Bytes(), *player)
				if err != nil {
					log.Println("Error in pill decrypt, err -", err)
				}

				if opcodes != nil {
					for _, v := range opcodes {
						pillEncrypt, err := pillInit.Encrypt(pillInit.SetOpcode(v).GetOutcomingCrumb())
						if err != nil {
							log.Println("Error in pill encrypt, err -", err)
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
