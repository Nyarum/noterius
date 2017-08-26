package entities

import (
	"time"

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
	Player *actor.PID
}

type DeletePlayer struct {
	Player *actor.PID
}

type GlobalTick struct {
	Now time.Time
}
