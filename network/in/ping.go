package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type Ping struct {
}

func (d Ping) Opcode() uint16 {
	return common.OP_CLIENT_PING
}

func (d *Ping) Unpack(pr *barrel.Processor) {
}
