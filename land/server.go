package land

import (
	"github.com/Nyarum/noterius/core"

	"log"
	"net"
)

// Application struct for project and his variables
type Application struct {
	Config core.Config
}

// Run function for starting server
func (a *Application) Run() (err error) {
	listen, err := net.Listen("tcp", a.Config.IP+":"+a.Config.Port)
	if err != nil {
		return
	}

	for {
		_, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
