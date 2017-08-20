package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type CreateCharacter struct {
	Name    string
	Map     string
	LenLook uint16
	Look    common.CharacterLookSub
}

func (c CreateCharacter) Opcode() uint16 {
	return common.OP_CLIENT_NEWCHA
}

func (c *CreateCharacter) Unpack(pr *barrel.Processor) {
	pr.ReadString1251(&c.Name)
	pr.ReadString1251(&c.Map)

	pr.ReadUint16(&c.LenLook)
	pr.ReadUint8(&c.Look.SynType)
	pr.ReadUint16(&c.Look.Race)
	pr.ReadUint8(&c.Look.BoatCheck)
}
