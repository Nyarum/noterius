package models

import kallax "gopkg.in/src-d/go-kallax.v1"

type Player struct {
	kallax.Model `table:"players" pk:"id"`
	ID           kallax.ULID
	Username     string
	Email        string
	Password     string
}
