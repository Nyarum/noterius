package pill

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/interface"
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/pill/incoming"
	"github.com/Nyarum/noterius/pill/outcoming"

	"errors"
)

type Pill struct {
	incomingCrumbs  map[int]interfaces.PillDecoder
	outcomingCrumbs map[int]interfaces.PillEncoder

	opcode int
}

func NewPill() *Pill {
	return &Pill{
		incomingCrumbs: map[int]interfaces.PillDecoder{
			431: &incoming.AuthCrumb{},
			432: &incoming.ExitCrumb{},
		},
		outcomingCrumbs: map[int]interfaces.PillEncoder{
			931: &outcoming.CharacterListCrumb{},
			940: &outcoming.DateCrumb{},
		},
	}
}

func (p *Pill) SetOpcode(opcode int) *Pill {
	p.opcode = opcode

	return p
}

func (p *Pill) GetIncomingCrumb() interfaces.PillDecoder {
	return p.incomingCrumbs[p.opcode]
}

func (p *Pill) GetOutcomingCrumb() interfaces.PillEncoder {
	return p.outcomingCrumbs[p.opcode]
}

func (p *Pill) Encrypt(pe interfaces.PillEncoder) ([]byte, error) {
	netes := network.NewParser([]byte{})

	data := pe.Process().PostHandler(netes)
	netes.Reset()

	header := Header{Len: uint16(len(data) + 8), UniqueId: 128, Opcode: uint16(p.opcode)}

	netes.SetEndian(network.LittleEndian).WriteUint16(header.Len)
	netes.SetEndian(network.BigEndian).WriteUint32(header.UniqueId)
	netes.SetEndian(network.LittleEndian).WriteUint16(header.Opcode)
	netes.WriteBytes([]byte(data))

	err := netes.Error()
	if err != nil {
		return nil, err
	}

	return netes.Bytes(), nil
}

func (p *Pill) Decrypt(buf []byte, player entitie.Player) ([]int, error) {
	var (
		header Header          = Header{}
		netes  *network.Parser = network.NewParser(buf)
	)

	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Len)
	netes.SetEndian(network.BigEndian).ReadUint32(&header.UniqueId)
	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Opcode)

	crumb := p.SetOpcode(int(header.Opcode)).GetIncomingCrumb()
	if crumb == nil {
		return nil, errors.New("Crumb is not found")
	}

	crumbProcess, err := crumb.PreHandler(netes).Process(player)
	if err != nil {
		return nil, err
	}

	err = netes.Error()
	if err != nil {
		return nil, err
	}

	return crumbProcess, nil
}
