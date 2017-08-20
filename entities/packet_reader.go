package entities

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/noterius/network/in"
	"go.uber.org/zap"
)

type PacketReader struct {
	Player       *actor.PID
	World        *actor.PID
	PacketSender *actor.PID
	Logger       *zap.SugaredLogger
}

func (state *PacketReader) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *in.Auth:
		state.Player.Tell(msg)
	case *in.Exit:
		state.PacketSender.Tell(Logout{})
	case *in.Ping:
	case *in.NewSecret:
		state.Player.Tell(msg)
	case *in.ChangeSecret:
		state.Player.Tell(msg)
	case *in.DeleteCharacter:
		state.Player.Tell(msg)
	case *in.CreateCharacter:
		state.Player.Tell(msg)
	}
}
