package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type DeleteCharacter struct {
	ErrorCode uint16
}

func (d DeleteCharacter) Opcode() uint16 {
	return common.OP_SERVER_DELCHA
}

func (d *DeleteCharacter) Pack(pr *barrel.Processor) {
	pr.WriteUint16(d.ErrorCode)
}
