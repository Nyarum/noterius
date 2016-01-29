package packet

import (
	"fmt"
	"time"

	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
)

func init() {
	Register(431, (*IncomingAuth)(&IncomingAuth{}).Packet)
	Register(432, (*IncomingExit)(&IncomingExit{}).Packet)
	Register(931, (*OutcomingCharacters)(&OutcomingCharacters{}).Packet)
	Register(940, (*OutcomingDate)(&OutcomingDate{}).Packet)
}

type OutcomingDate struct {
	Time string
}

func (i *OutcomingDate) Packet() (func(netes network.Netes), func(player *entitie.Player)) {
	handler := func(netes network.Netes) {
		netes.WriteString(i.Time)

		return
	}

	process := func(player *entitie.Player) {
		timeNow := time.Now()
		i.Time = fmt.Sprintf("[%d-%d %d:%d:%d:%d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000)

		return
	}

	return handler, process
}

type IncomingAuth struct {
	Key           string
	Login         string
	Password      string
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
}

func (i *IncomingAuth) Packet() (func(netes network.Netes), func(player *entitie.Player)) {
	handler := func(netes network.Netes) {
		netes.ReadString(&i.Key)
		netes.ReadString(&i.Login)
		netes.ReadString(&i.Password)
		netes.ReadString(&i.MAC)
		netes.ReadUint16(&i.IsCheat)
		netes.ReadUint16(&i.ClientVersion)

		return
	}

	process := func(player *entitie.Player) {
		player.Stats.Name = i.Login

		return
	}

	return handler, process
}

type OutcomingCharacters struct {
	ErrorCode  uint16
	Key        []byte
	Flag       uint8
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
}

func (i *OutcomingCharacters) Packet() (func(netes network.Netes), func(player *entitie.Player)) {
	handler := func(netes network.Netes) {
		netes.WriteUint16(uint16(0))
		netes.WriteBytes([]byte{0x00, 0x08, 0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49})
		netes.WriteUint8(uint8(0))
		netes.WriteUint8(uint8(1))
		netes.WriteUint32(uint32(0))
		netes.WriteUint32(uint32(12820))

		return
	}

	process := func(player *entitie.Player) {
		return
	}

	return handler, process
}

type IncomingExit struct {
}

func (i *IncomingExit) Packet() (func(netes network.Netes), func(player *entitie.Player)) {
	handler := func(netes network.Netes) {
		return
	}

	process := func(player *entitie.Player) {
		player.Buffers.GetEC() <- struct{}{}

		return
	}

	return handler, process
}
