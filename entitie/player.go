package entitie

import (
	"time"

	"github.com/Nyarum/noterius/core"
)

type Player struct {
	Login     string
	Character []Character
	CreatedAt time.Time
	UpdatedAt time.Time
	Buffers   *core.Buffers
}

func NewPlayer(buffers *core.Buffers) *Player {
	return &Player{Buffers: buffers}
}
