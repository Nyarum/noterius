package entities

import (
	"net"

	"go.uber.org/zap"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/noterius/network"
)

type PacketSender struct {
	Client          net.Conn
	Network         network.INetwork
	Logger          *zap.SugaredLogger
	CloseConnection chan struct{}
}

func (state *PacketSender) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case SendPacket:
		buf, err := state.Network.Marshal(msg.Packet)
		if err != nil {
			state.Logger.Errorw("Error marshal packet", "error", err)
			return
		}

		ln, err := state.Client.Write(buf)
		if err != nil {
			state.Logger.Errorw("Error send packet", "error", err)
			return
		}

		state.Logger.Debugw("Packet has sent", "len", ln)
	case SendPacketWithLogout:
		buf, err := state.Network.Marshal(msg.Packet)
		if err != nil {
			state.Logger.Errorw("Error marshal packet", "error", err)
			return
		}

		ln, err := state.Client.Write(buf)
		if err != nil {
			state.Logger.Errorw("Error send packet", "error", err)
			return
		}

		state.Logger.Debugw("Packet has sent", "len", ln)

		context.Self().Tell(Logout{})
	case Logout:
		state.Client.Close()
		state.CloseConnection <- struct{}{}
	}
}
