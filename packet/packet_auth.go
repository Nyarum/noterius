package packet

import (
	"fmt"
	"strings"
	"time"

	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/crypt"
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/support"
)

type OutcomingDate struct {
	Time string
}

func (i *OutcomingDate) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.WriteString(i.Time)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		timeNow := time.Now()
		i.Time = fmt.Sprintf("[%d-%d %d:%d:%d:%d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000)

		player.Time = i.Time

		return nil
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

func (i *IncomingAuth) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.ReadString(&i.Key)
		netes.ReadString(&i.Login)
		netes.ReadString(&i.Password)
		netes.ReadString(&i.MAC)
		netes.ReadUint16(&i.IsCheat)
		netes.ReadUint16(&i.ClientVersion)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		user, err := db.User().GetByName(i.Login)
		if err != nil {
			player.Error = &support.PlayerIsNotFound

			return []int{support.OP_SERVER_LOGIN}
		}

		encryptPassword, err := crypt.EncryptPassword(strings.ToUpper(user.Password[:24]), player.Time)
		if err != nil {
			player.Error = &support.CustomError{0, err}

			return []int{support.OP_SERVER_LOGIN}
		}

		if encryptPassword != i.Password {
			player.Error = &support.PasswordIncorrect

			return []int{support.OP_SERVER_LOGIN}
		}

		if user.IsActive {
			player.Error = &support.PlayerInGame

			return []int{support.OP_SERVER_LOGIN}
		}

		user.IsActive = true
		err = user.Update(user.ID)
		if err != nil {
			player.Error = &support.CustomError{0, err}

			return []int{support.OP_SERVER_LOGIN}
		}

		player.ID = user.ID

		return []int{support.OP_SERVER_LOGIN}
	}

	return handler, process
}

type CharacterLook struct {
	SynType   uint8
	Race      uint16
	BoatCheck uint8
	Items     [10]struct {
		Id   uint16
		Pass [160]byte
	}
	Hair uint16
}

type Character struct {
	Flag  uint8
	Name  string
	Job   string
	Level uint16
	Look  CharacterLook
}

type OutcomingAuth struct {
	ErrorCode  uint16
	Key        []byte
	Characters []Character
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
}

func (i *OutcomingAuth) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.WriteUint16(i.ErrorCode)
		netes.WriteBytes(i.Key)

		netes.WriteUint8(uint8(len(i.Characters)))
		for _, character := range i.Characters {
			netes.WriteUint8(character.Flag)
			netes.WriteString(character.Name)
			netes.WriteString(character.Job)
			netes.WriteUint16(character.Level)

			netes.WriteUint16(uint16(1626))
			netes.WriteUint8(character.Look.SynType)
			netes.WriteUint16(character.Look.Race)
			netes.WriteUint8(character.Look.BoatCheck)

			for _, item := range character.Look.Items {
				netes.WriteUint16(item.Id)
				netes.WriteBytes(item.Pass[:])
				/*
					netes.WriteUint16(item.Num)
					netes.WriteUint16(item.Endure)
					netes.WriteUint16(item.MaxEndure)
					netes.WriteUint16(item.Energy)
					netes.WriteUint16(item.MaxEnergy)
					netes.WriteUint8(item.ForgeLv)
					netes.WriteBool(item.Valid)
					netes.WriteUint8(item.CheckNext1)
					netes.WriteUint32(item.DbParam[0])
					netes.WriteUint32(item.DbParam[1])
					netes.WriteUint8(item.CheckNext2)

					for _, attr := range item.Attrs {
						netes.WriteUint16(attr.Id)
						netes.WriteUint16(attr.Value)
					}
				*/
			}

			netes.WriteUint16(character.Look.Hair)
		}

		netes.WriteUint8(i.Pincode)
		netes.WriteUint32(i.Encryption)
		netes.WriteUint32(i.DwFlag)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			switch player.Error {
			case &support.PlayerInGame:
				i.ErrorCode = player.Error.Code
			case &support.PlayerIsNotFound:
				i.ErrorCode = player.Error.Code
			case &support.PasswordIncorrect:
				i.ErrorCode = player.Error.Code
			default:
				i.ErrorCode = player.Error.Code
			}

			return nil
		}

		i.ErrorCode = 0
		i.Key = []byte{0x00, 0x08, 0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49}
		i.Pincode = 1
		i.Encryption = 0
		i.DwFlag = 12820

		// Only test data
		for b := 0; b < 3; b++ {
			character := Character{
				Flag:  1,
				Name:  "Haruki",
				Job:   "golang-ru.slack.com",
				Level: 1000,
			}
			character.Look.Race = 806

			i.Characters = append(i.Characters, character)
		}

		return nil
	}

	return handler, process
}

type IncomingExit struct {
}

func (i *IncomingExit) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		return
	}

	process := func(player *entitie.Player) []int {
		user := db.User()
		user.IsActive = false
		err := user.Update(player.ID)
		if err != nil {
			player.Error = &support.CustomError{0, err}
		}

		player.Buffers.GetEC() <- struct{}{}

		return nil
	}

	return handler, process
}
