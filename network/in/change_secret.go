package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type ChangeSecret struct {
	PasswordOld string
	PasswordNew string
}

func (c ChangeSecret) Opcode() uint16 {
	return common.OP_CLIENT_UPDATE_PASSWORD2
}

func (c *ChangeSecret) Unpack(pr *barrel.Processor) {
	pr.ReadString(&c.PasswordOld)
	pr.ReadString(&c.PasswordNew)
}
