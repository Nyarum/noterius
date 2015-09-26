package station

import (
	"github.com/Nyarum/noterius/core"

	"log"
	"time"
)

type DatabaseRobot struct {
}

func NewDatabaseRobot() *DatabaseRobot {
	return &DatabaseRobot{}
}

func (d *DatabaseRobot) ContiniusParse() {

}

func (d *DatabaseRobot) SaveOnTimeout(config core.Config) {
	for {
		select {
		case <-time.After(time.Duration(config.Database.TimeoutSave) * time.Minute):
			log.Println("Test timeout for database")
		}
	}
}
