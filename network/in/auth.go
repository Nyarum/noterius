package in

import (
	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/opcodes"
)

type Auth struct {
	Key           string
	Login         string
	Password      string
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
}

func (a Auth) Opcode() uint16 {
	return opcodes.OP_CLIENT_LOGIN
}

func (a *Auth) Unpack(pr *barrel.Processor) {
	pr.ReadString(&a.Key)
	pr.ReadString(&a.Login)
	pr.ReadString(&a.Password)
	pr.ReadString(&a.MAC)
	pr.ReadUint16(&a.IsCheat)
	pr.ReadUint16(&a.ClientVersion)
}
