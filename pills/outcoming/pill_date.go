package outcoming

import (
	"github.com/Nyarum/noterius/library/network"
)

type PillDate struct {
	Time string
}

func (pd *PillDate) Process() {
	timeNow := time.Now()
	pd.Time = fmt.Sprintf("[%d-%d %d:%d:%d:%d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000)
}

func (pd *PillDate) Encrypt(netes network.Netes) {
	netes.WriteString(pd.Time)
}
