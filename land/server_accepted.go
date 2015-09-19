package land

import (
	"github.com/Nyarum/noterius/core"

	"bytes"
	"log"
)

// ClientLive method for accept new connection from socket
func ClientLive(buffers Buffers, conf core.Config) {
	buffer := bytes.NewBuffer([]byte{})
	for getBytes := range buffers.GetReadChannel() {
		buffer.Write(getBytes)

		if conf.Base.Test {
			panic("Client is break :D")
		} else {
			log.Printf("Print message from client: % x", bytes.TrimRight(buffer.Bytes(), string([]byte{0x00})))
			buffers.GetWriteChannel() <- []byte{0x00, 0x02}
		}

		buffer.Reset()
	}
}
