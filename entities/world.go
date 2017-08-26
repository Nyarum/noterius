package entities

import (
	"database/sql"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type World struct {
	DB      *sql.DB
	Players *actor.PIDSet
}

func NewWorld(db *sql.DB) *World {
	return &World{
		DB:      db,
		Players: actor.NewPIDSet(),
	}
}

func (state *World) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case GlobalTick:
		state.Players.ForEach(func(i int, pid actor.PID) {
			pid.Tell(msg)
		})
	case AddPlayer:
		state.Players.Add(msg.Player)
	case DeletePlayer:
		state.Players.Remove(msg.Player)
	}
}

func (state *World) InWorld(context actor.Context) {

}
