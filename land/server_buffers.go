package land

import (
	"github.com/Nyarum/noterius/core"

	"io"
	"log"
	"net"
)

type Buffers struct {
	WriteChannel chan string
	ReadChannel  chan string
}

func NewBuffers() *Buffers {
	return &Buffers{
		WriteChannel: make(chan string),
		ReadChannel:  make(chan string),
	}
}

func (b *Buffers) GetWriteChannel() chan string {
	return b.WriteChannel
}

func (b *Buffers) GetReadChannel() chan string {
	return b.ReadChannel
}

func (b *Buffers) WriteHandler(c net.Conn) {
	for v := range b.WriteChannel {
		c.Write([]byte(v))
	}
}

func (b *Buffers) ReadHandler(c net.Conn, conf core.Config) {
	var (
		bytesAlloc []byte = make([]byte, conf.Option.LenBuffer)
	)

	for {
		if _, err := c.Read(bytesAlloc); err == io.EOF {
			log.Printf("Client [%v] is disconnected\n", c.RemoteAddr())
			return
		} else if err != nil {
			log.Printf("Client [%v] is error read packet, err - %v\n", c.RemoteAddr(), err)
		}

		b.ReadChannel <- string(bytesAlloc)
	}
}
