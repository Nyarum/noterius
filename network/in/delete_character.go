package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type DeleteCharacter struct {
	Name   string
	Secret string
}

func (d DeleteCharacter) Opcode() uint16 {
	return common.OP_CLIENT_DELCHA
}

func (d *DeleteCharacter) Unpack(pr *barrel.Processor) {
	pr.ReadString(&d.Name)
	pr.ReadString(&d.Secret)
}
