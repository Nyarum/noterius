package out

import "github.com/Nyarum/barrel"

type IOut interface {
	Opcode() uint16
	Pack(pr *barrel.Processor)
}

type Out struct {
}
