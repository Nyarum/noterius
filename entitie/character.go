package entitie

type Character struct {
	Stats    Stats
	Position Position
}

func NewCharacter() *Character {
	return &Character{}
}
