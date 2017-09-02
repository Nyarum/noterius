package out

import (
	"encoding/binary"
	"fmt"

	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type Auth struct {
	ErrorCode uint16
	// ErrorText string (if ErrorCode != 0)
	LenKey     uint16
	Key        []byte
	Characters []common.CharacterSub
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
}

func (a *Auth) SetPincode(pincode *string) {
	if pincode != nil && len(*pincode) == 32 {
		a.Pincode = 1
	}
}

func (d Auth) Opcode() uint16 {
	return common.OP_SERVER_LOGIN
}

func (d *Auth) Pack(pr *barrel.Processor) {
	pr.WriteUint16(d.ErrorCode)

	if d.ErrorCode == 0 {
		// Static key
		d.Key = []byte{0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49}

		pr.WriteUint16(uint16(len(d.Key)))
		pr.WriteBytes(d.Key)

		pr.WriteUint8(uint8(len(d.Characters)))
		for _, character := range d.Characters {
			pr.WriteUint8(character.Flag)

			if character.Flag == 1 {
				pr.WriteString1251(character.Name)
				pr.WriteString1251(character.Job)
				pr.WriteUint16(character.Level)

				fmt.Println(uint16(binary.Size(character.Look)))
				pr.WriteUint16(uint16(binary.Size(character.Look))) // Statically size look of character
				character.Look.Write(pr)
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
		character := common.CharacterSub{
			Flag:  1,
			Name:  "Haruki",
			Job:   "golang-ru.slack.com",
			Level: 1000,
		}
		character.Look.TypeID = 2

		d.Characters = append(d.Characters, character)
	}

	return d
}
