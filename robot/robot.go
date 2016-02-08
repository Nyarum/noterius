package robot

import "github.com/Nyarum/noterius/core"

type RobotFactory interface {
	Process(core.Config)
}

type Robot struct {
	Factories []RobotFactory
}

func NewRobot() *Robot {
	return &Robot{
		Factories: []RobotFactory{
			NewDatabase(),
		},
	}
}
