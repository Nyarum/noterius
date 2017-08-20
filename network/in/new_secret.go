package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type NewSecret struct {
	Password string
}

func (n NewSecret) Opcode() uint16 {
	return common.OP_CLIENT_CREATE_PASSWORD2
}

func (n *NewSecret) Unpack(pr *barrel.Processor) {
	pr.ReadString(&n.Password)
}
