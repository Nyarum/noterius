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
	}
}

func (state *World) InWorld(context actor.Context) {

}
