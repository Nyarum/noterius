package entities

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

type Player struct {
	World        *actor.PID
	PacketSender *actor.PID
}

func (state *Player) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Logout:
		// Something we do with database and other services
		// and exit

		state.PacketSender.Tell(msg)
	}
}
