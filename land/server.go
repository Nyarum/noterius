package land

import (
	"github.com/Nyarum/noterius/core"

	"log"
	"net"
)

// Application struct for project and his variables
type Application struct {
	Config       core.Config
	ErrorHandler func(c net.Conn)
}

// Run function for starting server
func (a *Application) Run() (err error) {
	listen, err := net.Listen("tcp", a.Config.IP+":"+a.Config.Port)
	if err != nil {
		return
	}

	for {
		client, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func(c net.Conn, test bool) {
			defer a.ErrorHandler(c)

			log.Println("Client is connected:", c.RemoteAddr())

			for {
				if test {
					panic("Client is break :D")
				} else {
					log.Println("Work is good")
					break
				}
			}
		}(client, a.Config.IsTest)
	}
}
