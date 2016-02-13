package packet

import (
	"errors"

	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/manager"
)

type PacketFactory func(*manager.Manager) (func(network.Netes), func(*entitie.Player) []int)

type Packet struct {
	pills   map[int]PacketFactory
	Player  *entitie.Player
	Manager *manager.Manager
}

type PacketHeader struct {
	Len      uint16
	UniqueId uint32
	Opcode   uint16
}

func NewPacket(player *entitie.Player, manager *manager.Manager) *Packet {
	return &Packet{
		pills: map[int]PacketFactory{
			431: (*IncomingAuth)(&IncomingAuth{}).Packet,
			432: (*IncomingExit)(&IncomingExit{}).Packet,
			931: (*OutcomingCharacters)(&OutcomingCharacters{}).Packet,
			940: (*OutcomingDate)(&OutcomingDate{}).Packet,
		},
		Player: player,
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

	handler, process := pck(p.Manager)

	process(p.Player)
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

func (p *Packet) Decode(buf []byte) ([]int, error) {
	var (
		header PacketHeader    = PacketHeader{}
		netes  *network.Parser = network.NewParser(buf)
	)

	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Len)
	netes.SetEndian(network.BigEndian).ReadUint32(&header.UniqueId)
	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Opcode)

	pck, err := p.GetPck(int(header.Opcode))
	if err != nil {
		return nil, err
	}

	handler, process := pck(p.Manager)

	handler(netes)
	opcodes := process(p.Player)

	err = netes.Error()
	if err != nil {
		return nil, err
	}

	return opcodes, nil
}
