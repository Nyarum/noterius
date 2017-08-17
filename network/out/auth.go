package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/opcodes"
)

type CharacterLookSub struct {
	SynType   uint8
	Race      uint16
	BoatCheck uint8
	Items     [10]struct {
		Id   uint16
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

type Auth struct {
	ErrorCode  uint16
	Key        []byte
	Characters []CharacterSub
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
}

func (a *Auth) SetPincode(pincode *uint16) {
	if pincode != nil {
		a.Pincode = 1
	}
}

func (d Auth) Opcode() uint16 {
	return opcodes.OP_SERVER_LOGIN
}

func (d *Auth) Pack(pr *barrel.Processor) {
	pr.WriteUint16(d.ErrorCode)

	if d.ErrorCode == 0 {
		// Static key
		d.Key = []byte{0x00, 0x08, 0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49}
		pr.WriteBytes(d.Key)

		pr.WriteUint8(uint8(len(d.Characters)))
		for _, character := range d.Characters {
			pr.WriteUint8(character.Flag)

			if character.Flag == 1 {
				pr.WriteString1251(character.Name)
				pr.WriteString1251(character.Job)
				pr.WriteUint16(character.Level)

				pr.WriteUint16(uint16(1626))
				pr.WriteUint8(character.Look.SynType)
				pr.WriteUint16(character.Look.Race)
				pr.WriteUint8(character.Look.BoatCheck)

				for _, item := range character.Look.Items {
					pr.WriteUint16(item.Id)
					pr.WriteBytes(item.Pass[:])
					/*
						pr.WriteUint16(item.Num)
						pr.WriteUint16(item.Endure)
						pr.WriteUint16(item.MaxEndure)
						pr.WriteUint16(item.Energy)
						pr.WriteUint16(item.MaxEnergy)
						pr.WriteUint8(item.ForgeLv)
						pr.WriteBool(item.Valid)
						pr.WriteUint8(item.CheckNext1)
						pr.WriteUint32(item.DbParam[0])
						pr.WriteUint32(item.DbParam[1])
						pr.WriteUint8(item.CheckNext2)
						for _, attr := range item.Attrs {
							pr.WriteUint16(attr.Id)
							pr.WriteUint16(attr.Value)
						}
					*/
				}

				pr.WriteUint16(character.Look.Hair)
			}
		}

		pr.WriteUint8(d.Pincode)
		pr.WriteUint32(d.Encryption)
		pr.WriteUint32(12820)
	}
}

func (d *Auth) SetTestData() *Auth {
	d.Pincode = 1

	// Only test data
	for b := 0; b < 3; b++ {
		character := CharacterSub{
			Flag:  1,
			Name:  "Haruki",
			Job:   "golang-ru.slack.com",
			Level: 1000,
		}
		character.Look.Race = 806

		d.Characters = append(d.Characters, character)
	}

	return d
}
