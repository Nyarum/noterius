package entities

type Monster struct {
	Stats    Stats
	Position Position
}

func NewMonster() *Monster {
	return &Monster{}
}
