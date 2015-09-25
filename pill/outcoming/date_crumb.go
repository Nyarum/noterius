package outcoming

import (
	"github.com/Nyarum/noterius/interface"
	"github.com/Nyarum/noterius/library/network"

	"fmt"
	"time"
)

type DateCrumb struct {
	Time string
}

func (d *DateCrumb) Process() interfaces.PillEncoder {
	timeNow := time.Now()
	d.Time = fmt.Sprintf("[%d-%d %d:%d:%d:%d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000)

	return d
}

func (d *DateCrumb) PostHandler(netes network.Netes) string {
	netes.WriteString(d.Time)

	return string(netes.Bytes())
}
