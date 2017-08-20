package common

type CharacterLookSub struct {
	SynType   uint8
	Race      uint16
	BoatCheck uint8
	Items     [10]struct {
		ID   uint16
		Pass [160]byte
	}
	Hair uint16
}

type CharacterSub struct {
	Flag  uint8
	Name  string
	Job   string
	Level uint16
	Look  CharacterLookSub
}

func (c *CharacterSub) SetFlag(flag bool) {
	if flag {
		c.Flag = 1
	}
}
