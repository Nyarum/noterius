package land

import (
	"github.com/Nyarum/noterius/core"

	"bytes"
	"encoding/binary"
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

// WriteHandler function for write bytes to socket in loop from channel
func (b *Buffers) WriteHandler(c net.Conn) {
	for v := range b.WriteChannel {
		c.Write([]byte(v))
	}
}

// ReadHandler function for read bytes from socket in loop to channel
func (b *Buffers) ReadHandler(c net.Conn, conf core.Config) {
	var (
		bytesAlloc []byte = make([]byte, conf.Option.LenBuffer)
	)

	buf := bytes.NewBuffer(bytesAlloc)
	for {
		_, err := c.Read(bytesAlloc)
		if err == io.EOF {
			log.Printf("Client [%v] is disconnected\n", c.RemoteAddr())
			return
		} else if err != nil {
			log.Printf("Client [%v] is error read packet, err - %v\n", c.RemoteAddr(), err)
		}

		lenFromData := int(binary.LittleEndian.Uint16(buf.Bytes()[0:2]))
		if buf.Len() < lenFromData {
			continue
		}

		b.ReadChannel <- string(buf.Next(lenFromData))
	}
}
