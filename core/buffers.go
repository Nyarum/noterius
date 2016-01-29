package core

import (
	log "github.com/Sirupsen/logrus"

	"bytes"
	"encoding/binary"
	"io"
	"net"
	"time"
)

// Buffers struct for read and write channels
type Buffers struct {
	writeChannel chan []byte
	readChannel  chan []byte
	exitChannel  chan struct{}
}

// NewBuffers method for init Buffers struct
func NewBuffers() *Buffers {
	return &Buffers{
		writeChannel: make(chan []byte),
		readChannel:  make(chan []byte),
		exitChannel:  make(chan struct{}),
	}
}

// GetWC method for get writeChannel from Buffers struct
func (b *Buffers) GetWC() chan []byte {
	return b.writeChannel
}

// GetRC method for get readChannel from Buffers struct
func (b *Buffers) GetRC() chan []byte {
	return b.readChannel
}

// GetEC method for get exitChannel from Buffers struct
func (b *Buffers) GetEC() chan struct{} {
	return b.exitChannel
}

func (b *Buffers) Close() {
	close(b.readChannel)
	close(b.writeChannel)
	close(b.exitChannel)

	return
}

// WriteHandler method for write bytes to socket in loop from channel
func (b *Buffers) WriteHandler(c net.Conn) {
	defer ErrorNetworkHandler(c)

	for v := range b.writeChannel {
		c.Write(v)
	}
}

// ReadHandler method for read bytes from socket in loop to channel
func (b *Buffers) ReadHandler(c net.Conn, conf Config) {
	var (
		buf     *bytes.Buffer = bytes.NewBuffer([]byte{})
		tempBuf []byte        = make([]byte, 2048)
	)

	// Monitoring close connection
	go func() {
		for v := range b.GetEC() {
			var _ struct{} = v
			c.Close()
		}
	}()

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
			b.readChannel <- sendBuffer
		}
	}
}
