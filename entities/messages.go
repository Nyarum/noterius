package entities

import (
	"github.com/AsynkronIT/protoactor-go/actor"
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

type RecordTime struct {
	Time string
}

type AddPlayer struct {
	ID     int64
	Player *actor.PID
}

type DeletePlayer struct {
	ID int64
}
