package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type NewSecret struct {
	ErrorCode uint16
}

func (n NewSecret) Opcode() uint16 {
	return common.OP_SERVER_UPDATE_PASSWORD2
}

func (n *NewSecret) Pack(pr *barrel.Processor) {
	pr.WriteUint16(n.ErrorCode)
}
