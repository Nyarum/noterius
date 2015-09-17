package outcoming

import (
	"github.com/Nyarum/noterius/interfaces"
	"github.com/Nyarum/noterius/library/network"

	"fmt"
	"time"
)

type CrumbDate struct {
	Time string
}

func (cd *CrumbDate) Process() interfaces.PillEncoder {
	timeNow := time.Now()
	cd.Time = fmt.Sprintf("[%d-%d %d:%d:%d:%d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000)

	return cd
}

func (cd *CrumbDate) PostHandler(netes network.Netes) string {
	netes.WriteString(cd.Time)

	return string(netes.Bytes())
}
