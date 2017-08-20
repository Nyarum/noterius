package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type ChangeSecret struct {
	ErrorCode uint16
}

func (c ChangeSecret) Opcode() uint16 {
	return common.OP_SERVER_UPDATE_PASSWORD2
}

func (c *ChangeSecret) Pack(pr *barrel.Processor) {
	pr.WriteUint16(c.ErrorCode)
}
