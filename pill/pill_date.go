package pill

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"

	"fmt"
	"time"
)

type Date struct {
	Time    string
	err     error
	opcodes []int
}

func (p *Date) Error() error {
	return p.err
}

func (p *Date) Opcodes() []int {
	return p.opcodes
}

func (p *Date) Handler(netes network.Netes) PillFactory {
	netes.WriteString(p.Time)

	return p
}

func (p *Date) Process(player *entitie.Player) PillFactory {
	timeNow := time.Now()
	p.Time = fmt.Sprintf("[%d-%d %d:%d:%d:%d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000)

	return p
}
