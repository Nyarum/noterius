package entities

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

type World struct {
	Players      map[int64]*actor.PID
	PacketSender *actor.PID
}

func (state *World) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	}
}

func (state *World) InWorld(context actor.Context) {

}
