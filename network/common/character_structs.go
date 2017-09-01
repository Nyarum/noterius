package common

import "github.com/Nyarum/barrel"

type InstAttribute struct {
	ID    uint16
	Value uint16
}

type Item struct {
	ID        uint16
	Num       uint16
	Endure    [2]uint16
	Energy    [2]uint16
	ForgeLv   uint8
	IsValid   uint8
	Change    uint8
	DbParam   [2]uint32
	InstFlag  uint8
	InstAttrs [5]InstAttribute
	ItemAttrs [58]uint16
}

type CharacterLookSub struct {
	SynType     uint8
	TypeID      uint16
	UnknownFlag uint8
	Items       [10]Item
	Hair        uint16
}

func (c *CharacterLookSub) Read(pr *barrel.Processor) {
	pr.ReadUint8(&c.SynType)
	pr.ReadUint16(&c.TypeID)
	pr.ReadUint8(&c.UnknownFlag)
	pr.ReadUint16(&c.Hair)

	for k := range c.Items {
		pr.ReadUint16(&c.Items[k].ID)

		if c.Items[k].ID == 0 {
			continue
		}

		pr.ReadUint32(&c.Items[k].DBID)
		pr.ReadUint16(&c.Items[k].Num)

		for kEndure := range c.Items[k].Endure {
			pr.ReadUint16(&c.Items[k].Endure[kEndure])
		}

		for kEnergy := range c.Items[k].Energy {
			pr.ReadUint16(&c.Items[k].Energy[kEnergy])
		}

		pr.ReadUint8(&c.Items[k].ForgeLv)
		pr.ReadUint8(&c.Items[k].IsValid)
		pr.ReadUint8(&c.Items[k].Change)

		for kDbParam := range c.Items[k].DbParam {
			pr.ReadUint32(&c.Items[k].DbParam[kDbParam])
		}

		pr.ReadUint8(&c.Items[k].InstFlag)

		if c.Items[k].InstAttrs[0].ID > 0 {
			for kInstAttrs := range c.Items[k].InstAttrs {
				pr.ReadUint16(&c.Items[k].InstAttrs[kInstAttrs].ID)
				pr.ReadUint16(&c.Items[k].InstAttrs[kInstAttrs].Value)
			}
		}
	}
}

func (c *CharacterLookSub) Write(pr *barrel.Processor) {
	pr.WriteUint8(c.SynType)
	pr.WriteUint16(c.TypeID)
	pr.WriteUint8(c.UnknownFlag)
	pr.WriteUint16(c.Hair)

	for k := range c.Items {
		pr.WriteUint16(c.Items[k].ID)
		pr.WriteUint32(c.Items[k].DBID)

		if c.Items[k].ID == 0 {
			continue
		}

		pr.WriteUint16(c.Items[k].Num)

		for kEndure := range c.Items[k].Endure {
			pr.WriteUint16(c.Items[k].Endure[kEndure])
		}

		for kEnergy := range c.Items[k].Energy {
			pr.WriteUint16(c.Items[k].Energy[kEnergy])
		}

		pr.WriteUint8(c.Items[k].ForgeLv)
		pr.WriteUint8(c.Items[k].IsValid)
		pr.WriteUint8(1)

		for kDbParam := range c.Items[k].DbParam {
			pr.WriteUint32(c.Items[k].DbParam[kDbParam])
		}

		if c.Items[k].InstAttrs[0].ID > 0 {
			pr.WriteUint8(1)
			for kInstAttrs := range c.Items[k].InstAttrs {
				pr.WriteUint16(c.Items[k].InstAttrs[kInstAttrs].ID)
				pr.WriteUint16(c.Items[k].InstAttrs[kInstAttrs].Value)
			}
		} else {
			pr.WriteUint8(0)
		}
	}
}

type BoatLookSub struct {
	PosID     uint16
	BoatID    uint16
	Header    uint16
	Body      uint16
	Engine    uint16
	Cannon    uint16
	Equipment uint16
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
	Attrs     [5]InstAttribute
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
