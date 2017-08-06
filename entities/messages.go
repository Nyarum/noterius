package entities

import (
	"github.com/Nyarum/noterius/network/in"
	"github.com/Nyarum/noterius/network/out"
)

type ReadPacket struct {
	Len int
	Buf []byte
}

type SendPacket struct {
	Packet out.IOut
}

type SendPacketWithLogout struct {
	Packet out.IOut
}

type Logout struct {
}

type Auth struct {
	*in.Auth
}

type RecordTime struct {
	Time string
}
