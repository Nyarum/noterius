package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/opcodes"
)

type Ping struct {
}

func (d Ping) Opcode() uint16 {
	return opcodes.OP_CLIENT_PING
}

func (d *Ping) Unpack(pr *barrel.Processor) {
}
