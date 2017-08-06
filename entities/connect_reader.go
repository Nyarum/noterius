package entities

import (
	"encoding/binary"
	"net"

	"go.uber.org/zap"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/noterius/network"
)

type ConnectReader struct {
	Client       net.Conn
	PacketReader *actor.PID
	Network      network.INetwork
	Logger       *zap.SugaredLogger
}

func (state *ConnectReader) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case ReadPacket:
		if msg.Len == 2 {
			state.Client.Write([]byte{0x00, 0x02})
			return
		}

		uniqueID := binary.LittleEndian.Uint32(msg.Buf[2:6])
		opcode := binary.BigEndian.Uint16(msg.Buf[6:8])

		state.Logger.Debugw("Received a new packet", "len", msg.Len, "uniqueID", uniqueID, "opcode", opcode)

		if msg.Len >= 8 {
			packet, err := state.Network.Unmarshal(opcode, msg.Buf[8:])
			if err != nil {
				state.Logger.Errorw("Error unmarshal packet", "error", err)
				return
			}

			state.PacketReader.Tell(packet)
		}
	}
}
