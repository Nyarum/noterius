package land

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
			defer a.ErrorHandler(c)

			log.Println("Client is connected:", c.RemoteAddr())

			var (
				bytesAlloc      []byte      = make([]byte, conf.Option.LenBuffer)
				readBytesSocket chan string = make(chan string)
			)

			go ClientLive(c, readBytesSocket, conf)

			for {
				_, err := c.Read(bytesAlloc)
				if err != nil {
					log.Printf("Client [%v] is disconnect after read packet error, err - %v", c.RemoteAddr(), err)
					break
				}

				readBytesSocket <- string(bytesAlloc)
			}
		}(client, a.Config)
	}
}
