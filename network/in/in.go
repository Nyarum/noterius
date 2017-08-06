package in

import "github.com/Nyarum/barrel"

type IIn interface {
	Opcode() uint16
	Unpack(pr *barrel.Processor)
}

type In struct {
}
