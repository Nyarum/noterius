package common

import "github.com/Nyarum/barrel"

type InstAttribute struct {
	ID    uint16
	Value uint16
}

type Item struct {
	ID         uint16
	Num        uint16
	Endure     [2]uint16
	Energy     [2]uint16
	ForgeLv    uint8
	PassValue  uint8
	DbParam    [2]uint32
	InstAttrs  [5]InstAttribute
	ItemAttrs  [58]uint16
	InitFlag   uint8
	PassValue2 uint8
	Valid      bool
	Change     bool
}

type CharacterLookSub struct {
	Ver    uint16
	TypeID uint16
	Items  [10]Item
	Hair   uint16
}

func (c *CharacterLookSub) Read(pr *barrel.Processor) {
	pr.SetEndian(barrel.LittleEndian)
	pr.ReadUint16(&c.Ver)
	pr.ReadUint16(&c.TypeID)

	for k := range c.Items {
		pr.ReadUint16(&c.Items[k].ID)
		pr.ReadUint16(&c.Items[k].Num)

		for kEndure := range c.Items[k].Endure {
			pr.ReadUint16(&c.Items[k].Endure[kEndure])
		}

		for kEnergy := range c.Items[k].Energy {
			pr.ReadUint16(&c.Items[k].Energy[kEnergy])
		}

		pr.ReadUint8(&c.Items[k].ForgeLv)
		pr.ReadUint8(&c.Items[k].PassValue)

		for kDbParam := range c.Items[k].DbParam {
			pr.ReadUint32(&c.Items[k].DbParam[kDbParam])
		}

		for kInstAttrs := range c.Items[k].InstAttrs {
			pr.ReadUint16(&c.Items[k].InstAttrs[kInstAttrs].ID)
			pr.ReadUint16(&c.Items[k].InstAttrs[kInstAttrs].Value)
		}

		for kItemAttrs := range c.Items[k].ItemAttrs {
			pr.ReadUint16(&c.Items[k].ItemAttrs[kItemAttrs])
		}

		pr.ReadUint8(&c.Items[k].InitFlag)
		pr.ReadUint8(&c.Items[k].PassValue2)
		pr.ReadBool(&c.Items[k].Valid)
		pr.ReadBool(&c.Items[k].Change)
	}

	pr.ReadUint16(&c.Hair)
	pr.SetEndian(barrel.BigEndian)
}

func (c *CharacterLookSub) Write(pr *barrel.Processor) {
	pr.SetEndian(barrel.LittleEndian)
	pr.WriteUint16(c.Ver)
	pr.WriteUint16(c.TypeID)

	for k := range c.Items {
		pr.WriteUint16(c.Items[k].ID)
		pr.WriteUint16(c.Items[k].Num)

		for kEndure := range c.Items[k].Endure {
			pr.WriteUint16(c.Items[k].Endure[kEndure])
		}

		for kEnergy := range c.Items[k].Energy {
			pr.WriteUint16(c.Items[k].Energy[kEnergy])
		}

		pr.WriteUint8(c.Items[k].ForgeLv)
		pr.WriteUint8(c.Items[k].PassValue)

		for kDbParam := range c.Items[k].DbParam {
			pr.WriteUint32(c.Items[k].DbParam[kDbParam])
		}

		for kInstAttrs := range c.Items[k].InstAttrs {
			pr.WriteUint16(c.Items[k].InstAttrs[kInstAttrs].ID)
			pr.WriteUint16(c.Items[k].InstAttrs[kInstAttrs].Value)
		}

		for kItemAttrs := range c.Items[k].ItemAttrs {
			pr.WriteUint16(c.Items[k].ItemAttrs[kItemAttrs])
		}

		pr.WriteUint8(c.Items[k].InitFlag)
		pr.WriteUint8(c.Items[k].PassValue2)
		pr.WriteBool(c.Items[k].Valid)
		pr.WriteBool(c.Items[k].Change)
	}

	pr.WriteUint16(c.Hair)
	pr.SetEndian(barrel.BigEndian)
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
