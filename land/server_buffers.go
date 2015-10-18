package land

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/pill"
	log "github.com/Sirupsen/logrus"

	"bytes"
	"encoding/binary"
	"io"
	"net"
	"time"
)

// Buffers struct for read and write channels
type Buffers struct {
	WriteChannel chan []byte
	ReadChannel  chan []byte
}

// NewBuffers method for init Buffers struct
func NewBuffers() *Buffers {
	return &Buffers{
		WriteChannel: make(chan []byte),
		ReadChannel:  make(chan []byte),
	}
}

// GetWriteChannel method for get WriteChannel from Buffers struct
func (b *Buffers) GetWriteChannel() chan []byte {
	return b.WriteChannel
}

// GetReadChannel method for get ReadChannel from Buffers struct
func (b *Buffers) GetReadChannel() chan []byte {
	return b.ReadChannel
}

// WriteHandler method for write bytes to socket in loop from channel
func (b *Buffers) WriteHandler(c net.Conn) {
	defer core.ErrorNetworkHandler(c)

	// Write one packet for client with time.Now()
	pillInit := pill.NewPill()
	packet, err := pillInit.Encrypt(pillInit.SetOpcode(940).GetOutcomingCrumb())
	if err != nil {
		log.WithError(err).Error("Error in pill encrypt")
	}

	c.Write(packet)

	for v := range b.WriteChannel {
		c.Write(v)
	}
}

// ReadHandler method for read bytes from socket in loop to channel
func (b *Buffers) ReadHandler(c net.Conn, conf core.Config) {
	var (
		buf     *bytes.Buffer = bytes.NewBuffer([]byte{})
		tempBuf []byte        = make([]byte, 2048)
	)

	for {
		ln, err := c.Read(tempBuf)
		if err != nil {
			if err.(net.Error).Timeout() {
				log.WithField("address", c.RemoteAddr()).Warn("Client is timeout")
			}

			if err == io.EOF {
				log.WithField("address", c.RemoteAddr()).Warn("Client is disconnected")
			}

			log.WithField("address", c.RemoteAddr()).Warn("Client has closed connection")

			return
		}

		c.SetReadDeadline(time.Now().Add(10 * time.Second))
		buf.Write(tempBuf[:ln])

		var lastGotLen int
		readLen := func() bool {
			if buf.Len() < 2 {
				return false
			}

			lastGotLen = int(binary.BigEndian.Uint16(buf.Bytes()[0:2]))
			if buf.Len() < lastGotLen {
				return false
			}

			return true
		}

		for readLen() {
			sendBuffer := make([]byte, ln)
			copy(sendBuffer, buf.Next(lastGotLen))
			b.ReadChannel <- sendBuffer
		}
	}
}
