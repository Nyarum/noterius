package entitie

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/support"
)

type Player struct {
	ID       int64
	Stats    Stats
	Position Position

	Buffers *core.Buffers
	Error   *support.CustomError
	Time    string
}

func NewPlayer(buffers *core.Buffers) *Player {
	return &Player{Buffers: buffers}
}
