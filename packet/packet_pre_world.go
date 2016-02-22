package packet

import (
	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/support"
)

type IncomingBeginPlay struct {
	Name string
}

func (i *IncomingBeginPlay) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.ReadString(&i.Name)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		return []int{support.OP_SERVER_ENTERMAP}
	}

	return handler, process
}

type MapCharacterBag struct {
	GridId   uint16
	GridItem MapCharacterItem
}

type MapCharacterState struct {
	Id uint8
	Lv uint8
}

type MapCharacterSkill struct {
	Id         uint16
	State      uint8
	Lv         uint8
	UseSp      uint16
	UseEndure  uint16
	UseEnergy  uint16
	ResumeTime uint32
	Range      [4][2]byte
}

type MapCharacterItem struct {
	Id         uint16
	Num        uint16
	Endure     uint16
	MaxEndure  uint16
	Energy     uint16
	MaxEnergy  uint16
	ForgeLv    uint8
	Valid      bool
	CheckNext1 uint8
	DbParam    [2]uint32
	CheckNext2 uint8
	Attrs      [5]struct {
		Id    uint16
		Value uint16
	}
}

type MapCharacterLook struct {
	SynType   uint8
	Race      uint16
	BoatCheck uint8
	HairId    uint16
	Items     [10]MapCharacterItem
}

type MapEvent struct {
	EnityId   uint32
	EnityType uint8
	EventId   uint16
	EventName string
}

type MapCharacterBase struct {
	WorldId      uint32
	CharCId      uint32
	CharMId      uint32
	CharCName    string
	Unknown      uint16
	GmLv         uint8
	Handle       uint16
	CtrlType     uint8
	CharMName    string
	MottoName    string
	Icon         uint16
	Unknown2     uint16
	GuildId      uint16
	GuildName    string
	GuildMotto   string
	StallName    string
	State        uint16
	PosX         uint32
	PosY         uint32
	PosRadius    uint32
	PosAngle     uint16
	TeamLeaderId uint32
}

type OutcomingEnterMap struct {
	EnterRet            uint16
	AutoLock            uint8
	KitbagLock          uint8
	EnterType           uint8
	IsNew               bool
	Map                 string
	IsTeam              bool
	MapCharacterBase    MapCharacterBase
	MapSide             uint8
	MapEvent            MapEvent
	MapCharacterLook    MapCharacterLook
	IsPK                bool
	MapCharacterAppends [4]struct {
		LookId uint16
		Valid  bool
	}
	MapCharacterSkills struct {
		Id     uint16
		Type   uint8
		Num    uint16
		Skills []MapCharacterSkill
	}
	MapCharacterStates struct {
		Num    uint8
		States []MapCharacterState
	}
	MapOptions struct {
		Type    uint8
		Num     uint16
		Options [74]struct {
			AttrId uint8
			Value  uint32
		}
	}
	MapBags struct {
		Type      uint8
		KeybagNum uint16
		Bags      []MapCharacterBag
	}
	MapVisible   uint16
	MapShortcuts struct {
		Shortcuts [36]struct {
			Type   uint8
			GridId uint16
		}
	}
	BoatNum     uint8
	WorldMainId uint32
}

func (i *OutcomingEnterMap) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.WriteUint16(i.EnterRet)
		netes.WriteUint8(i.AutoLock)
		netes.WriteUint8(i.KitbagLock)
		netes.WriteUint8(i.EnterType)
		netes.WriteBool(i.IsNew)
		netes.WriteString(i.Map)
		netes.WriteBool(i.IsTeam)

		netes.WriteUint32(i.MapCharacterBase.WorldId)
		netes.WriteUint32(i.MapCharacterBase.CharCId)
		netes.WriteUint32(i.MapCharacterBase.CharMId)
		netes.WriteString(i.MapCharacterBase.CharCName)
		netes.WriteUint16(i.MapCharacterBase.Unknown)
		netes.WriteUint8(i.MapCharacterBase.GmLv)
		netes.WriteUint16(i.MapCharacterBase.Handle)
		netes.WriteUint8(i.MapCharacterBase.CtrlType)
		netes.WriteString(i.MapCharacterBase.CharMName)
		netes.WriteString(i.MapCharacterBase.MottoName)
		netes.WriteUint16(i.MapCharacterBase.Icon)
		netes.WriteUint16(i.MapCharacterBase.Unknown2)
		netes.WriteUint16(i.MapCharacterBase.GuildId)
		netes.WriteString(i.MapCharacterBase.GuildName)
		netes.WriteString(i.MapCharacterBase.GuildMotto)
		netes.WriteString(i.MapCharacterBase.StallName)
		netes.WriteUint16(i.MapCharacterBase.State)
		netes.WriteUint32(i.MapCharacterBase.PosX)
		netes.WriteUint32(i.MapCharacterBase.PosY)
		netes.WriteUint32(i.MapCharacterBase.PosRadius)
		netes.WriteUint16(i.MapCharacterBase.PosAngle)
		netes.WriteUint32(i.MapCharacterBase.TeamLeaderId)

		netes.WriteUint8(i.MapSide)

		netes.WriteUint32(i.MapEvent.EnityId)
		netes.WriteUint8(i.MapEvent.EnityType)
		netes.WriteUint16(i.MapEvent.EventId)
		netes.WriteString(i.MapEvent.EventName)

		netes.WriteUint8(i.MapCharacterLook.SynType)
		netes.WriteUint16(i.MapCharacterLook.Race)
		netes.WriteUint8(i.MapCharacterLook.BoatCheck)
		netes.WriteUint16(i.MapCharacterLook.HairId)
		for _, item := range i.MapCharacterLook.Items {
			netes.WriteUint16(item.Id)
			if item.Id != 0 {
				netes.WriteUint16(item.Num)
				netes.WriteUint16(item.Endure)
				netes.WriteUint16(item.MaxEndure)
				netes.WriteUint16(item.Energy)
				netes.WriteUint16(item.MaxEnergy)
				netes.WriteUint8(item.ForgeLv)
				netes.WriteBool(item.Valid)
				netes.WriteUint8(item.CheckNext1)
				netes.WriteUint32(item.DbParam[0])
				netes.WriteUint32(item.DbParam[1])
				netes.WriteUint8(item.CheckNext2)
				for _, attr := range item.Attrs {
					netes.WriteUint16(attr.Id)
					netes.WriteUint16(attr.Value)
				}
			}
		}

		netes.WriteBool(i.IsPK)

		for _, v := range i.MapCharacterAppends {
			netes.WriteUint16(v.LookId)

			if v.LookId != 0 {
				netes.WriteBool(v.Valid)
			}
		}

		netes.WriteUint16(i.MapCharacterSkills.Id)
		netes.WriteUint8(i.MapCharacterSkills.Type)
		netes.WriteUint16(i.MapCharacterSkills.Num)
		for _, v := range i.MapCharacterSkills.Skills {
			netes.WriteUint16(v.Id)
			netes.WriteUint8(v.State)
			netes.WriteUint8(v.Lv)
			netes.WriteUint16(v.UseSp)
			netes.WriteUint16(v.UseEndure)
			netes.WriteUint16(v.UseEnergy)
			netes.WriteUint32(v.ResumeTime)
			for _, rangeItem := range v.Range {
				netes.WriteBytes(rangeItem[:])
			}
		}

		netes.WriteUint8(i.MapCharacterStates.Num)
		for _, v := range i.MapCharacterStates.States {
			netes.WriteUint8(v.Id)
			netes.WriteUint8(v.Lv)
		}

		netes.WriteUint8(i.MapOptions.Type)
		netes.WriteUint16(i.MapOptions.Num)
		for _, v := range i.MapOptions.Options {
			netes.WriteUint8(v.AttrId)
			netes.WriteUint32(v.Value)
		}

		netes.WriteUint8(i.MapBags.Type)
		netes.WriteUint16(i.MapBags.KeybagNum)
		for _, v := range i.MapBags.Bags {
			netes.WriteUint16(v.GridId)
			if v.GridItem.Id != 0 {
				netes.WriteUint16(v.GridItem.Id)
				netes.WriteUint16(v.GridItem.Num)
				netes.WriteUint16(v.GridItem.Endure)
				netes.WriteUint16(v.GridItem.MaxEndure)
				netes.WriteUint16(v.GridItem.Energy)
				netes.WriteUint16(v.GridItem.MaxEnergy)
				netes.WriteUint8(v.GridItem.ForgeLv)
				netes.WriteBool(v.GridItem.Valid)
				netes.WriteUint8(v.GridItem.CheckNext1)
				netes.WriteUint32(v.GridItem.DbParam[0])
				netes.WriteUint32(v.GridItem.DbParam[1])
				netes.WriteUint8(v.GridItem.CheckNext2)
				for _, attr := range v.GridItem.Attrs {
					netes.WriteUint16(attr.Id)
					netes.WriteUint16(attr.Value)
				}
			}
		}

		netes.WriteUint16(i.MapVisible)

		for _, v := range i.MapShortcuts.Shortcuts {
			netes.WriteUint8(v.Type)
			netes.WriteUint16(v.GridId)
		}

		netes.WriteUint8(i.BoatNum)
		netes.WriteUint32(i.WorldMainId)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		i.EnterType = 1
		i.IsNew = true
		i.Map = "garner"
		i.IsTeam = true
		i.MapCharacterBase = MapCharacterBase{4, 123, 123, "Pilotka", 25346, 0, 11437, 1, "Pilotka", "", 4, 0, 0, "My Guild", "", "", 1, 224700, 270400, 40, 0, 0}
		i.MapEvent = MapEvent{123, 1, 0, ""}
		i.MapCharacterLook = MapCharacterLook{0, 1, 0, 2291, [10]MapCharacterItem{}}
		i.MapVisible = 65535

		i.MapOptions.Num = 74

		mapCharacterBags := []MapCharacterBag{}
		for n := 0; n <= 23; n++ {
			mapCharacterBags = append(mapCharacterBags, MapCharacterBag{uint16(n), MapCharacterItem{Id: 0}})
		}
		i.MapBags.Type = 0
		i.MapBags.KeybagNum = uint16(len(mapCharacterBags))

		for _, v := range i.MapShortcuts.Shortcuts {
			v.Type = 255
			v.GridId = 65280
		}

		i.WorldMainId = 123

		return nil
	}

	return handler, process
}
