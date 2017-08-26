package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type BagTempSync struct {
}

func (b BagTempSync) Opcode() uint16 {
	return common.OP_CLIENT_KITBAGTEMP_SYNC
}

func (b *BagTempSync) Unpack(pr *barrel.Processor) {

}
