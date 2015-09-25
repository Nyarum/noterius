package entitie

import (
	"net"
)

type Player struct {
	Stats    Stats
	Position Position

	Connection net.Conn
}

func NewPlayer(c net.Conn) *Player {
	return &Player{Connection: c}
}
