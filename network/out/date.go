package out

import (
	"fmt"
	"time"

	"github.com/Nyarum/barrel"
	"github.com/Nyarum/noterius/network/opcodes"
)

type Date struct {
	Time string
}

func (d Date) Opcode() uint16 {
	return opcodes.OP_SERVER_CHAPSTR
}

func (d *Date) Pack(pr *barrel.Processor) {
	pr.WriteString(d.Time)
}

func (d *Date) GetCurrentTime() string {
	timeNow := time.Now()

	return fmt.Sprintf("[%02d-%02d %02d:%02d:%02d:%03d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000)
}
