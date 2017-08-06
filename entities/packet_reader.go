package entities

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/noterius/network/in"
	"github.com/Nyarum/noterius/network/out"
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
		state.Logger.Debugw("Received Auth packet", "details", msg)

		state.PacketSender.Tell(SendPacket{
			Packet: (&out.Auth{}).SetTestData(),
		})
	case *in.Exit:
		state.Logger.Debugw("Received Exit packet", "details", msg)

		state.Player.Tell(Logout{})
	case *in.Ping:
		state.Logger.Debugw("Received Ping packet", "details", msg)
	}
}
