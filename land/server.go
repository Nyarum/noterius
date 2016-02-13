package land

import (
	"github.com/Nyarum/noterius/core"
	log "github.com/Sirupsen/logrus"

	"net"
)

// Application struct for project and his variables
type Application struct {
	Config       core.Config
	DatabaseInfo core.DatabaseInfo
}

// Run method for starting server
func (a *Application) Run() (err error) {
	listen, err := net.Listen("tcp", a.Config.Base.IP+":"+a.Config.Base.Port)
	if err != nil {
		return
	}

	for {
		client, err := listen.Accept()
		if err != nil {
			log.WithError(err).Error("Error in accept connection")
			continue
		}

		go func(c net.Conn, conf core.Config) {
			var buffers *core.Buffers = core.NewBuffers()

			defer func() {
				buffers.Close()
				core.ErrorNetworkHandler(c)
			}()

			log.WithFields(log.Fields{
				"address": c.RemoteAddr(),
			}).Info("Client is connected")

			go ConnectHandler(buffers, *a, c)
			go buffers.WriteHandler(c)

			buffers.ReadHandler(c, conf)
		}(client, a.Config)
	}
}
