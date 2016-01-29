package packet

import (
	"errors"

	"github.com/Nyarum/noterius/library/network"
)

type PacketFactory func() (func(netes network.Netes), func())

var packetFuncs map[int]PacketFactory = make(map[int]PacketFactory)

func Register(opcode int, funcCall PacketFactory) {
	packetFuncs[opcode] = funcCall
}

type Packet struct {
	pills map[int]PacketFactory
}

type PacketHeader struct {
	Len      uint16
	UniqueId uint32
	Opcode   uint16
}

func NewPacket() *Packet {
	return &Packet{
		pills: packetFuncs,
	}
}

func (p *Packet) GetPck(opcode int) (PacketFactory, error) {
	get, ok := p.pills[opcode]
	if !ok {
		return nil, errors.New("Pill is not found")
	}

	return get, nil
}

func (p *Packet) Encode(opcode int) ([]byte, error) {
	netes := network.NewParser([]byte{})

	pck, err := p.GetPck(opcode)
	if err != nil {
		return nil, err
	}

	handler, process := pck()

	process()
	handler(netes)

	data := string(netes.Bytes())
	netes.Reset()

	header := PacketHeader{Len: uint16(len(data) + 8), UniqueId: 128, Opcode: uint16(opcode)}

	netes.SetEndian(network.LittleEndian).WriteUint16(header.Len)
	netes.SetEndian(network.BigEndian).WriteUint32(header.UniqueId)
	netes.SetEndian(network.LittleEndian).WriteUint16(header.Opcode)
	netes.WriteBytes([]byte(data))

	err = netes.Error()
	if err != nil {
		return nil, err
	}

	return netes.Bytes(), nil
}

func (p *Packet) Decode(buf []byte) error {
	var (
		header PacketHeader    = PacketHeader{}
		netes  *network.Parser = network.NewParser(buf)
	)

	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Len)
	netes.SetEndian(network.BigEndian).ReadUint32(&header.UniqueId)
	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Opcode)

	pck, err := p.GetPck(int(header.Opcode))
	if err != nil {
		return err
	}

	handler, process := pck()

	handler(netes)
	process()

	err = netes.Error()
	if err != nil {
		return err
	}

	return nil
}
