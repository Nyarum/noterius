package network

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/in"
	"github.com/Nyarum/noterius/network/out"
)

type Error string

func (n Error) Error() string {
	return string(n)
}

const (
	NotSupportedOpcode Error = "This opcode is not supported by the server"
)

type INetwork interface {
	Unmarshal(opcode uint16, b []byte) (in.IIn, error)
	Marshal(packet out.IOut) ([]byte, error)
}

type Network struct {
	packets   map[uint16]in.IIn
	processor *barrel.Processor
}

func NewNetwork() *Network {
	packets := make(map[uint16]in.IIn)
	packets[(in.Ping{}).Opcode()] = &in.Ping{}
	packets[(in.Auth{}).Opcode()] = &in.Auth{}
	packets[(in.Exit{}).Opcode()] = &in.Exit{}

	return &Network{
		packets:   packets,
		processor: barrel.NewProcessor([]byte{}),
	}
}

func (n *Network) Unmarshal(opcode uint16, b []byte) (in.IIn, error) {
	n.processor.WriteBytes(b)
	defer n.processor.Reset()

	packet, ok := n.packets[opcode]
	if !ok {
		return nil, NotSupportedOpcode
	}

	packet.Unpack(n.processor)

	if n.processor.Error() != nil {
		return nil, n.processor.Error()
	}

	return packet, nil
}

func (n *Network) Marshal(packet out.IOut) ([]byte, error) {
	packet.Pack(n.processor)
	defer n.processor.Reset()

	dataBuf := n.processor.Clone()
	n.processor.Reset()

	//r := rand.New().RandomSeed()

	n.processor.SetEndian(barrel.BigEndian).WriteUint16(uint16(len(dataBuf) + 8))
	n.processor.SetEndian(barrel.LittleEndian).WriteUint32(128)
	n.processor.SetEndian(barrel.BigEndian).WriteUint16(packet.Opcode())
	n.processor.WriteBytes(dataBuf)

	if n.processor.Error() != nil {
		return nil, n.processor.Error()
	}

	return n.processor.Bytes(), nil
}
