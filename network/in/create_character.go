package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type CreateCharacter struct {
	Name string
	Map  string
}

func (c CreateCharacter) Opcode() uint16 {
	return common.OP_CLIENT_NEWCHA
}

func (c *CreateCharacter) Unpack(pr *barrel.Processor) {

}
