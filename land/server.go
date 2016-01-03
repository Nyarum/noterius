package land

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/robot"
	log "github.com/Sirupsen/logrus"

	"net"
)

// Application struct for project and his variables
type Application struct {
	Config   core.Config
	Database core.Database
}

// Run method for starting server
func (a *Application) Run() (err error) {
	listen, err := net.Listen("tcp", a.Config.Base.IP+":"+a.Config.Base.Port)
	if err != nil {
		return
	}

	robot := robot.NewRobot()
	for _, factory := range robot.Factories {
		go factory.Process(a.Config)
	}

	for {
		client, err := listen.Accept()
		if err != nil {
			log.WithError(err).Error("Error in accept connection")
			continue
		}

		go func(c net.Conn, conf core.Config) {
			var buffers *Buffers = NewBuffers()

			defer func() {
				close(buffers.GetReadChannel())
				close(buffers.GetWriteChannel())
				core.ErrorNetworkHandler(c)
			}()

			log.WithFields(log.Fields{
				"address": c.RemoteAddr(),
			}).Info("Client is connected")

			go ClientLive(*buffers, conf, c)
			go buffers.WriteHandler(c)

			buffers.ReadHandler(c, conf)
		}(client, a.Config)
	}
}
