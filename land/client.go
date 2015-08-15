package land

import (
	"github.com/Nyarum/noterius/core"

	"bytes"
	"log"
	"net"
)

func ClientLive(c net.Conn, chReadBytes chan string, conf core.Config) {
	var (
		getBytes   string
		bytesAlloc []byte        = make([]byte, conf.Option.LenBuffer)
		buffer     *bytes.Buffer = bytes.NewBuffer(bytesAlloc)
	)

	for getBytes = range chReadBytes {
		copy(bytesAlloc, getBytes)

		if conf.Base.Test {
			panic("Client is break :D")
		} else {
			log.Printf("Print message from client: %v", string(buffer.Bytes()))
		}
	}
}
