package entitie

type NPC struct {
	Stats    Stats
	Position Position
}

func NewNPC() *NPC {
	return &NPC{}
}
