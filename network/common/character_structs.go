package common

type ItemAttribute struct {
	ID    uint16
	Value uint16
}

type Item struct {
	ID            uint16
	Num           uint16
	Durability    uint16
	MaxDurability uint16
	Energy        uint16
	MaxEnergy     uint16
	ForgeLv       uint8
	Valid         bool
	DbParam1      uint32
	DbParam2      uint32
	Attrs         [5]ItemAttribute
	Unknown2      [119]byte
	Unknown3      uint8
}

type CharacterLookSub struct {
	SynType   uint8
	Race      uint16
	BoatCheck uint8
	Items     [10]Item
	Hair      uint16
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

type CharacterInfo struct {
	WorldID    uint32
	CharCId    uint32
	CharMId    uint32
	CharCName  string
	Unknown    uint16
	GmLv       uint8
	Handle     uint16
	CtrlType   uint8
	CharMName  string
	MottoName  string
	Icon       uint16
	Unknown2   uint16
	GuildID    uint16
	GuildName  string
	GuildMotto string
	StallName  string
	State      uint16
	PosX       uint32
	PosY       uint32
	PosRadius  uint32
	PosAngle   uint16
	TeamLeadID uint32
}

type Bag struct {
	GridID    uint16
	ItemID    uint16
	Num       uint16
	Endure    uint16
	MaxEndure uint16
	Energy    uint16
	MaxEnergy uint16
	ForgeLv   uint8
	Valid     bool
	DbParam0  uint32
	DbParam1  uint32
	CheckNext uint8
	Attrs     [5]ItemAttribute
}

type CharacterBags struct {
	Type      uint8
	KeybagNum uint16
	Bags      []Bag
}

type CharacterSkills struct {
	ID     uint16
	Type   uint8
	Num    uint16
	Skills []Skill
}

type CharacterAppendLook struct {
	LookID uint16
	Valid  bool
}

type CharacterSkillStates struct {
	Num    uint8
	States []SkillState
}

type Attribute struct {
	ID    uint8
	Value uint32
}

type CharacterAttibutes struct {
	Type       uint8
	Num        uint16
	Attributes [74]Attribute
}

type Shortcut struct {
	Type   uint8
	GridID uint16
}
