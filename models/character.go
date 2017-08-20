package models

import kallax "gopkg.in/src-d/go-kallax.v1"

type Character struct {
	kallax.Model `table:"characters" pk:"id,autoincr"`
	kallax.Timestamps
	ID      int64
	Player  *Player `fk:"player_id,inverse"`
	Name    string  `unique:"true"`
	Job     string
	Level   uint16
	Race    uint16
	Enabled bool
}
