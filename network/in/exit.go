package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/opcodes"
)

type Exit struct {
}

func (a Exit) Opcode() uint16 {
	return opcodes.OP_CLIENT_LOGOUT
}

func (a *Exit) Unpack(pr *barrel.Processor) {

}
