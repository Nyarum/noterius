package land

import (
	"github.com/Nyarum/noterius/core"

	"log"
	"net"
	"time"
)

// Application struct for project and his variables
type Application struct {
	Config   core.Config
	Database core.Database
	Error    *core.Error
}

// Run method for starting server
func (a *Application) Run() (err error) {
	listen, err := net.Listen("tcp", a.Config.Base.IP+":"+a.Config.Base.Port)
	if err != nil {
		return
	}

	for {
		client, err := listen.Accept()
		client.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			log.Println(err)
			continue
		}

		go func(c net.Conn, conf core.Config) {
			var buffers *Buffers = NewBuffers()
			defer func() {
				close(buffers.GetReadChannel())
				close(buffers.GetWriteChannel())
				a.Error.NetworkHandler(c)
			}()

			log.Println("Client is connected:", c.RemoteAddr())

			go ClientLive(*buffers, conf)
			go buffers.WriteHandler(c)

			buffers.ReadHandler(c, conf)
		}(client, a.Config)
	}
}
