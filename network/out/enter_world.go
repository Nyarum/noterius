package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type EnterWorld struct {
	EnterRet    uint16
	AutoLock    uint8
	KitbagLock  uint8
	EnterType   uint8
	NewChar     uint8
	MapName     string
	CanTeam     uint8
	Info        common.CharacterInfo
	CharSide    uint8
	EntityEvent common.EntityEvent
	Look        common.CharacterLookSub
	IsPK        uint8
	AppendLooks [4]common.CharacterAppendLook
	Skills      common.CharacterSkills
	SkillStates common.CharacterSkillStates
	Attributes  common.CharacterAttibutes
	Bags        common.CharacterBags
	MapVisible  uint16
	Shortcuts   [36]common.Shortcut
	BoatNum     uint8
	CharMainID  uint32
}

func (c EnterWorld) Opcode() uint16 {
	return common.OP_SERVER_ENTERMAP
}

func (c *EnterWorld) Pack(pr *barrel.Processor) {
	pr.WriteUint16(c.EnterRet)
	pr.WriteUint8(c.AutoLock)
	pr.WriteUint8(c.KitbagLock)
	pr.WriteUint8(c.EnterType)
	pr.WriteUint8(c.NewChar)
	pr.WriteString1251(c.MapName)
	pr.WriteUint8(c.CanTeam)

	// Write info
	pr.WriteUint32(c.Info.WorldID)
	pr.WriteUint32(c.Info.CharCId)
	pr.WriteUint32(c.Info.CharMId)
	pr.WriteString1251(c.Info.CharCName)
	pr.WriteUint16(c.Info.Unknown)
	pr.WriteUint8(c.Info.GmLv)
	pr.WriteUint16(c.Info.Handle)
	pr.WriteUint8(c.Info.CtrlType)
	pr.WriteString1251(c.Info.CharMName)
	pr.WriteString1251(c.Info.MottoName)
	pr.WriteUint16(c.Info.Icon)
	pr.WriteUint16(c.Info.Unknown2)
	pr.WriteUint16(c.Info.GuildID)
	pr.WriteString1251(c.Info.GuildName)
	pr.WriteString1251(c.Info.GuildMotto)
	pr.WriteString1251(c.Info.StallName)
	pr.WriteUint16(c.Info.State)
	pr.WriteUint32(c.Info.PosX)
	pr.WriteUint32(c.Info.PosY)
	pr.WriteUint32(c.Info.PosRadius)
	pr.WriteUint16(c.Info.PosAngle)
	pr.WriteUint32(c.Info.TeamLeadID)

	pr.WriteUint8(c.CharSide)

	// Write entity event
	pr.WriteUint32(c.EntityEvent.EnityID)
	pr.WriteUint8(c.EntityEvent.EnityType)
	pr.WriteUint16(c.EntityEvent.EventID)
	pr.WriteString1251(c.EntityEvent.EventName)

	// Write look
	pr.WriteUint16(1626) // Statically size look of character
	pr.WriteUint8(c.Look.SynType)
	pr.WriteUint16(c.Look.Race)
	pr.WriteUint8(c.Look.BoatCheck)

	for _, item := range c.Look.Items {
		pr.WriteUint16(item.ID)
		pr.WriteUint16(item.Num)
		pr.WriteUint16(item.Durability)
		pr.WriteUint16(item.MaxDurability)
		pr.WriteUint16(item.Energy)
		pr.WriteUint16(item.MaxEnergy)
		pr.WriteUint8(item.ForgeLv)
		pr.WriteBool(item.Valid)
		pr.WriteUint32(item.DbParam1)
		pr.WriteUint32(item.DbParam2)
		for _, attr := range item.Attrs {
			pr.WriteUint16(attr.ID)
			pr.WriteUint16(attr.Value)
		}
		pr.WriteBytes(item.Unknown2[:])
		pr.WriteByte(item.Unknown3)
	}
	pr.WriteUint16(c.Look.Hair)

	pr.WriteUint8(c.IsPK)

	// Write append looks
	for _, v := range c.AppendLooks {
		pr.WriteUint16(v.LookID)

		if v.LookID != 0 {
			pr.WriteBool(v.Valid)
		}
	}

	// Write skills
	pr.WriteUint16(c.Skills.ID)
	pr.WriteUint8(c.Skills.Type)
	pr.WriteUint16(c.Skills.Num)

	if c.Skills.Num != 0 {
		for _, v := range c.Skills.Skills {
			pr.WriteUint16(v.ID)
			pr.WriteUint8(v.State)
			pr.WriteUint8(v.Lv)
			pr.WriteUint16(v.UseSp)
			pr.WriteUint16(v.UseEndure)
			pr.WriteUint16(v.UseEnergy)
			pr.WriteUint32(v.ResumeTime)

			for _, vR := range v.Range {
				for _, vR2 := range vR {
					pr.WriteByte(vR2)
				}
			}
		}
	}

	// Write skill states
	pr.WriteUint8(c.SkillStates.Num)

	if c.SkillStates.Num != 0 {
		for _, v := range c.SkillStates.States {
			pr.WriteUint8(v.ID)
			pr.WriteUint8(v.Lv)
		}
	}

	// Write attributes
	pr.WriteUint8(c.Attributes.Type)
	pr.WriteUint16(c.Attributes.Num)

	if c.Attributes.Num != 0 {
		for _, v := range c.Attributes.Attributes {
			pr.WriteUint8(v.ID)
			pr.WriteUint32(v.Value)
		}
	}

	// Write bags
	pr.WriteUint8(c.Bags.Type)

	if c.Bags.Type == 0 {
		pr.WriteUint16(c.Bags.KeybagNum)

		if c.Bags.KeybagNum != 0 {
			for _, v := range c.Bags.Bags {
				pr.WriteUint16(v.GridID)
				pr.WriteUint16(v.ItemID)

				if v.ItemID != 0 {
					pr.WriteUint16(v.Num)
					pr.WriteUint16(v.Endure)
					pr.WriteUint16(v.MaxEndure)
					pr.WriteUint16(v.Energy)
					pr.WriteUint16(v.MaxEnergy)
					pr.WriteUint8(v.ForgeLv)
					pr.WriteBool(v.Valid)
					pr.WriteUint32(v.DbParam0)
					pr.WriteUint32(v.DbParam1)
					pr.WriteUint8(v.CheckNext)
				} else if v.ItemID != 0 && v.CheckNext != 0 {
					for _, vA := range v.Attrs {
						pr.WriteUint16(vA.ID)
						pr.WriteUint16(vA.Value)
					}
				}
			}
		}
	}

	pr.WriteUint16(c.MapVisible)

	// Write shortcuts
	for _, v := range c.Shortcuts {
		pr.WriteUint8(v.Type)
		pr.WriteUint16(v.GridID)
	}

	pr.WriteUint8(c.BoatNum)
	pr.WriteUint32(c.CharMainID)
}

func (c *EnterWorld) SetTestChar() *EnterWorld {
	shortcut := common.Shortcut{255, 65280}
	shortcuts := [36]common.Shortcut{shortcut}
	attrb := common.ItemAttribute{0, 0}
	attrsb := [5]common.ItemAttribute{attrb, attrb, attrb, attrb, attrb}
	bag := common.Bag{0, 0, 0, 0, 0, 0, 0, 0, true, 0, 0, 0, attrsb}
	bags := []common.Bag{bag, bag, bag, bag}
	charkitbag := common.CharacterBags{0, 4, bags}
	attributes := [74]common.Attribute{}
	for i := 0; i < len(attributes); i++ {
		attributes[i] = common.Attribute{uint8(i), 1000}
	}

	charattributes := common.CharacterAttibutes{0, 74, attributes}
	states := []common.SkillState{}
	charskillstate := common.CharacterSkillStates{0, states}
	skills := []common.Skill{}
	charskillbag := common.CharacterSkills{0, 0, 0, skills}
	appendt := common.CharacterAppendLook{0, true}
	appends := [4]common.CharacterAppendLook{appendt, appendt, appendt, appendt}
	c.AppendLooks = appends
	c.IsPK = 0
	attr := common.ItemAttribute{0, 0}
	attrs := [5]common.ItemAttribute{attr, attr, attr, attr, attr}
	item := common.Item{0, 0, 0, 0, 0, 0, 0, true, 0, 0, attrs, [119]byte{}, 0}
	items := [10]common.Item{item, item, item, item, item, item, item, item, item, item}
	charlook := common.CharacterLookSub{0, 2, 0, items, 2291}
	entevent := common.EntityEvent{10263, 1, 0, ""}
	c.CharSide = 0

	charbase := common.CharacterInfo{2, 10263, 10263, "Свой парень", 25346, 0, 11437, 1, "Свой парень", "", 4, 0, 0, "Своя гильдия", "", "", 1, 217375, 278125, 40, 333, 0}

	c.EnterRet = 0
	c.AutoLock = 0
	c.KitbagLock = 0
	c.EnterType = 1
	c.NewChar = 0
	c.MapName = "garner"
	c.CanTeam = 1
	c.Info = charbase
	c.EntityEvent = entevent
	c.Look = charlook
	c.Skills = charskillbag
	c.SkillStates = charskillstate
	c.Attributes = charattributes
	c.Bags = charkitbag
	c.MapVisible = 65535
	c.Shortcuts = shortcuts
	c.BoatNum = 0
	c.CharMainID = 10263

	return c
}
