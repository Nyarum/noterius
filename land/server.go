package land

// First packet for client init connect with date
/*
	current := time.Now()
	date := fmt.Sprintf("[%d-%d %d:%d:%d:%d]", current.Month(), current.Day(), current.Hour(), current.Minute(), current.Second(), current.Nanosecond()/1000000)
	ln := uint16(11 + len(date))
	group := []byte{0x80, 0x00, 0x00, 0x00}
	opcode := uint16(940)

	netes := network.NewParser([]byte{})
	netes.Write(ln)
	netes.SetEndian(network.BigEndian).Write(group)
	netes.SetEndian(network.LittleEndian).Write(opcode)
	netes.SetEndian(network.LittleEndian).Write(date)

	c.Write(netes.Bytes())
*/

import (
	"github.com/Nyarum/noterius/core"

	"log"
	"net"
)

// Application struct for project and his variables
type Application struct {
	Config       core.Config
	Database     core.Database
	ErrorHandler func(c net.Conn)
}

// Run function for starting server
func (a *Application) Run() (err error) {
	listen, err := net.Listen("tcp", a.Config.Base.IP+":"+a.Config.Base.Port)
	if err != nil {
		return
	}

	for {
		client, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func(c net.Conn, conf core.Config) {
			var buffers *Buffers = NewBuffers()
			defer a.ErrorHandler(c)
			defer close(buffers.GetReadChannel())
			defer close(buffers.GetWriteChannel())

			log.Println("Client is connected:", c.RemoteAddr())

			go ClientLive(*buffers, conf)
			go buffers.WriteHandler(c)

			buffers.ReadHandler(c, conf)
		}(client, a.Config)
	}
}
