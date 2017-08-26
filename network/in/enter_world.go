package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type EnterWorld struct {
	Name string
}

func (e EnterWorld) Opcode() uint16 {
	return common.OP_CLIENT_BGNPLAY
}

func (e *EnterWorld) Unpack(pr *barrel.Processor) {
	pr.ReadString(&e.Name)
}
