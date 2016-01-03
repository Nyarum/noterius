package pill

import (
	"errors"

	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
)

type PillFactory interface {
	Error() error
	Opcodes() []int
	Handler(network.Netes) PillFactory
	Process(*entitie.Player) PillFactory
}

type Pill struct {
	packets map[int]PillFactory
	opcode  int
}

type Header struct {
	Len      uint16
	UniqueId uint32
	Opcode   uint16
}

func NewPill() *Pill {
	return &Pill{
		packets: map[int]PillFactory{
			AuthOpcode:          &Auth{},
			ExitOpcode:          &Exit{},
			CharacterListOpcode: &CharacterList{},
			DateOpcode:          &Date{},
		},
	}
}

func (p *Pill) GetPill(opcode int) PillFactory {
	p.opcode = opcode

	return p.packets[p.opcode]
}

func (p *Pill) Encrypt(pf PillFactory, player *entitie.Player) ([]byte, error) {
	netes := network.NewParser([]byte{})

	pf.Process(player).Handler(netes)
	data := string(netes.Bytes())
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

func (p *Pill) Decrypt(buf []byte, player *entitie.Player) ([]int, error) {
	var (
		header Header          = Header{}
		netes  *network.Parser = network.NewParser(buf)
	)

	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Len)
	netes.SetEndian(network.BigEndian).ReadUint32(&header.UniqueId)
	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Opcode)

	pill := p.GetPill(int(header.Opcode))
	if pill == nil {
		return nil, errors.New("Pill is not found")
	}

	process := pill.Handler(netes).Process(player)
	if process.Error() != nil {
		return nil, process.Error()
	}

	err := netes.Error()
	if err != nil {
		return nil, err
	}

	return process.Opcodes(), nil
}
