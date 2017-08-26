package entities

import (
	"database/sql"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type World struct {
	DB      *sql.DB
	Players map[int64]*actor.PID
}

func (state *World) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case AddPlayer:
		state.Players[msg.ID] = msg.Player
	case DeletePlayer:
		delete(state.Players, msg.ID)
	}
}

func (state *World) InWorld(context actor.Context) {

}
