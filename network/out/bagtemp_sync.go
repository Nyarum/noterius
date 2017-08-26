package out

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/common"
)

type BagTempSync struct {
	ErrorCode uint16
	Unknown   uint8
	Max       uint16
}

func (b BagTempSync) Opcode() uint16 {
	return common.OP_SERVER_KITBAGTEMP_SYNC
}

func (b *BagTempSync) Pack(pr *barrel.Processor) {
	pr.WriteUint16(b.ErrorCode)
	pr.WriteUint8(b.Unknown)
	pr.WriteUint16(b.Max)
}

func (b *BagTempSync) SetTestData() *BagTempSync {
	b.Unknown = 16
	b.Max = 65535

	return b
}
