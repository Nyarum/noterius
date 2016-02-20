package entitie

import "github.com/Nyarum/noterius/core"

type Player struct {
	ID       int64
	Stats    Stats
	Position Position

	Buffers *core.Buffers
	Error   error
	Time    string
}

func NewPlayer(buffers *core.Buffers) *Player {
	return &Player{Buffers: buffers}
}
