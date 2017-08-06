package models

import kallax "gopkg.in/src-d/go-kallax.v1"

type Player struct {
	kallax.Model `table:"players" pk:"id,autoincr"`
	kallax.Timestamps
	ID       int64
	Username string
	Email    string
	Password string
}
