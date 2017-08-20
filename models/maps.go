package models

import kallax "gopkg.in/src-d/go-kallax.v1"

type Map struct {
	kallax.Model `table:"maps" pk:"id,autoincr"`
	kallax.Timestamps
	ID   int64
	Name []string `unique:"true"`
}
