package outcoming

import (
	"github.com/Nyarum/noterius/interface"
	"github.com/Nyarum/noterius/library/network"
)

type CharacterListCrumb struct {
	Error      uint16
	Key        []byte
	Flag       uint8
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
}

func (cl *CharacterListCrumb) Process() interfaces.PillEncoder {
	return cl
}

func (cl *CharacterListCrumb) PostHandler(netes network.Netes) string {
	netes.WriteUint16(uint16(0))
	netes.WriteBytes([]byte{0x00, 0x08, 0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49})
	netes.WriteUint8(uint8(0))
	netes.WriteUint8(uint8(1))
	netes.WriteUint32(uint32(0))
	netes.WriteUint32(uint32(12820))

	return string(netes.Bytes())
}
