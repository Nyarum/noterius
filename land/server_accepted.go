package land

import (
	"github.com/Nyarum/noterius/core"

	"bytes"
	"log"
)

func ClientLive(buffers Buffers, conf core.Config) {
	var (
		bytesAlloc []byte        = make([]byte, conf.Option.LenBuffer)
		buffer     *bytes.Buffer = bytes.NewBuffer(bytesAlloc)
	)

	for getBytes := range buffers.GetReadChannel() {
		copy(bytesAlloc, getBytes)

		if conf.Base.Test {
			panic("Client is break :D")
		} else {
			log.Printf("Print message from client: % x", bytes.TrimRight(buffer.Bytes(), string([]byte{0x00})))
			buffers.GetWriteChannel() <- "Hello from server\n"
		}
	}
}
