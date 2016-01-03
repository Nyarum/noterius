package robot

import (
	"log"
	"time"

	"github.com/Nyarum/noterius/core"
)

type Database struct {
}

func NewDatabase() RobotFactory {
	return &Database{}
}

func (d *Database) Process(config core.Config) {
	for {
		select {
		case <-time.After(time.Duration(config.Database.TimeoutSave) * time.Minute):
			log.Println("Test timeout for database")
		}
	}
}
