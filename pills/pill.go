package pills

import (
	"github.com/Nyarum/noterius/library/network"
)

type PillEncoder interface {
	Encrypt()
}

type PillDecoder interface {
	Decrypt()
}

type Pill struct {
}

func (p *Pill) Encrypt() []byte {
	netes := network.NewParser([]byte{})
	netes.SetEndian(network.BigEndian).WriteUint16(0)
	netes.SetEndian(network.LittleEndian).WriteBytes([]byte{0x80, 0x00, 0x00, 0x00})
}

func (p *Pill) Decrypt() {

}
