package core

import (
	log "github.com/Sirupsen/logrus"

	"net"
	"os"
	"runtime/debug"
)

// ErrorNetworkHandle method for handler client accepted
func ErrorNetworkHandler(c net.Conn) {
	if r := recover(); r != nil {
		log.WithField("error", string(debug.Stack())).Error("Error in network")
	}

	c.Close()

}

// ErrorGlobalHandler method for handler global fatals
func ErrorGlobalHandler() {
	if r := recover(); r != nil {
		log.WithField("error", string(debug.Stack())).Error("Error in starting server")
		os.Exit(0)
	}
}
