package packet

import (
	"fmt"
	"strings"
	"time"

	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/crypt"
	"github.com/Nyarum/noterius/library/network"
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
			player.Error = PlayerIsNotFound

			return []int{OP_CHARACTERS}
		}

		encryptPassword, err := crypt.EncryptPassword(strings.ToUpper(user.Password[:24]), player.Time)
		if err != nil {
			player.Error = err

			return []int{OP_CHARACTERS}
		}

		if encryptPassword != i.Password {
			player.Error = PasswordIncorrect

			return []int{OP_CHARACTERS}
		}

		if user.IsActive {
			player.Error = PlayerInGame

			return []int{OP_CHARACTERS}
		}

		user.IsActive = true
		err = user.Update(user.ID)
		if err != nil {
			player.Error = err

			return []int{OP_CHARACTERS}
		}

		player.ID = user.ID

		return []int{OP_CHARACTERS}
	}

	return handler, process
}

type OutcomingCharacters struct {
	ErrorCode  uint16
	Key        []byte
	Flag       uint8
	Characters []struct {
		Flag  uint8
		Name  string
		Job   string
		Level uint16
		Look  struct {
			SynType   uint8
			Race      uint16
			BoatCheck uint8
			Items     [10]struct {
				Id         uint16
				Num        uint16
				Endure     uint16
				MaxEndure  uint16
				Energy     uint16
				MaxEnergy  uint16
				ForgeLv    uint8
				Valid      bool
				CheckNext1 uint8
				DbParam    [2]uint32
				CheckNext2 uint8
				Attrs      [5]struct {
					Id    uint16
					Value uint16
				}
			}
			Hair uint16
		}
	}
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
}

func (i *OutcomingCharacters) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.WriteUint16(i.ErrorCode)
		netes.WriteBytes(i.Key)
		netes.WriteUint8(i.Flag)
		netes.WriteUint8(i.Pincode)
		netes.WriteUint32(i.Encryption)
		netes.WriteUint32(i.DwFlag)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			switch player.Error {
			case PlayerInGame:
				i.ErrorCode = 1104
			case PlayerIsNotFound:
				i.ErrorCode = 1001
			case PasswordIncorrect:
				i.ErrorCode = 1002
			default:
				i.ErrorCode = 1000
			}

			return nil
		}

		i.ErrorCode = 0
		i.Key = []byte{0x00, 0x08, 0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49}
		i.Flag = 0
		i.Pincode = 1
		i.Encryption = 0
		i.DwFlag = 12820

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
			player.Error = err
		}

		player.Buffers.GetEC() <- struct{}{}

		return nil
	}

	return handler, process
}
