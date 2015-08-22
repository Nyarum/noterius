package pills

import (
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/pills/incoming"
	"github.com/Nyarum/noterius/pills/outcoming"
)

type PillEncoder interface {
	Process()
	PostHandler(network.Netes) string
}

type PillDecoder interface {
	PreHandler(network.Netes)
	Process()
}

type Pill struct {
	incomingCrumbs  map[int]PillDecoder
	outcomingCrumbs map[int]PillEncoder

	opcode int
}

func NewPill() *Pill {
	return &Pill{
		incomingCrumbs: map[int]PillDecoder{
			431: &incoming.CrumbAuth{},
		},
		outcomingCrumbs: map[int]PillEncoder{
			940: &outcoming.CrumbDate{},
		},
	}
}

func (p *Pill) SetOpcode(opcode int) *Pill {
	p.opcode = opcode

	return p
}

func (p *Pill) GetIncomingCrumb() PillDecoder {
	return p.incomingCrumbs[p.opcode]
}

func (p *Pill) GetOutcomingCrumb() PillEncoder {
	return p.outcomingCrumbs[p.opcode]
}

func (p *Pill) Encrypt(pe PillEncoder) []byte {
	netes := network.NewParser([]byte{})

	pe.Process()
	data := pe.PostHandler(netes)
	netes.Reset()

	netes.SetEndian(network.LittleEndian).WriteUint16(uint16(len(data) + 8))
	netes.SetEndian(network.LittleEndian).WriteBytes([]byte{0x80, 0x00, 0x00, 0x00})
	netes.SetEndian(network.LittleEndian).WriteUint16(uint16(p.opcode))
	netes.WriteBytes([]byte(data))

	return netes.Bytes()
}

func (p *Pill) Decrypt(pd PillDecoder, buf []byte) {
	network.NewParser(buf)
}
