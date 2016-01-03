package pill

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
)

type CharacterList struct {
	ErrorCode  uint16
	Key        []byte
	Flag       uint8
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
	err        error
	opcodes    []int
}

func (p *CharacterList) Error() error {
	return p.err
}

func (p *CharacterList) Opcodes() []int {
	return p.opcodes
}

func (p *CharacterList) Handler(netes network.Netes) PillFactory {
	netes.WriteUint16(uint16(0))
	netes.WriteBytes([]byte{0x00, 0x08, 0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49})
	netes.WriteUint8(uint8(0))
	netes.WriteUint8(uint8(1))
	netes.WriteUint32(uint32(0))
	netes.WriteUint32(uint32(12820))

	return p
}

func (p *CharacterList) Process(player *entitie.Player) PillFactory {
	return p
}
