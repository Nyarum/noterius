package entitie

import "github.com/Nyarum/noterius/core"

type Player struct {
	Stats    Stats
	Position Position

	Buffers *core.Buffers
}

func NewPlayer(buffers *core.Buffers) *Player {
	return &Player{Buffers: buffers}
}
