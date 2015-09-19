package entities

type Player struct {
	Stats    Stats
	Position Position
}

func NewPlayer() *Player {
	return &Player{}
}
