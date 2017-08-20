package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type Ping struct {
}

func (d Ping) Opcode() uint16 {
	return common.OP_SERVER_PING
}

func (d *Ping) Pack(pr *barrel.Processor) {
}
