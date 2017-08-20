package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type CreateCharacter struct {
	ErrorCode uint16
}

func (c CreateCharacter) Opcode() uint16 {
	return common.OP_SERVER_NEWCHA
}

func (c *CreateCharacter) Pack(pr *barrel.Processor) {
	pr.WriteUint16(c.ErrorCode)
}
