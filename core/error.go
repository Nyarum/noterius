package core

import (
	"log"
	"net"
	"os"
	"runtime/debug"
)

type Error struct{}

func NewError() *Error {
	return &Error{}
}

// NetworkHandler method for handler client accepted
func (e *Error) NetworkHandler(c net.Conn) {
	if r := recover(); r != nil {
		log.Printf("%s: %s\n", r, debug.Stack())
	}

	c.Close()
}

// GlobalHandler method for handler global fatals
func (e *Error) GlobalHandler() {
	if r := recover(); r != nil {
		log.Printf("%s: %s\n", r, debug.Stack())
		os.Exit(1)
	}
}
