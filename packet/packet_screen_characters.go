package packet

import (
	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/support"
	"github.com/aglyzov/charmap"
)

type IncomingDeleteCharacter struct {
	Name   string
	Secret string
}

func (i *IncomingDeleteCharacter) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.ReadString(&i.Name)
		netes.ReadString(&i.Secret)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		return []int{support.OP_SERVER_DELCHA}
	}

	return handler, process
}

type OutcomingDeleteCharacter struct {
	ErrorCode uint16
}

func (i *OutcomingDeleteCharacter) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.WriteUint16(i.ErrorCode)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		//i.ErrorCode = support.SecretPasswordIncorrect.Code

		return nil
	}

	return handler, process
}

type IncomingUpdatePassword struct {
	SecretOld string
	SecretNew string
}

func (i *IncomingUpdatePassword) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.ReadString(&i.SecretOld)
		netes.ReadString(&i.SecretNew)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		return []int{support.OP_SERVER_UPDATE_PASSWORD2}
	}

	return handler, process
}

type OutcomingUpdatePassword struct {
	ErrorCode uint16
}

func (i *OutcomingUpdatePassword) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.WriteUint16(i.ErrorCode)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		//i.ErrorCode = support.SecretPasswordIncorrect.Code

		return nil
	}

	return handler, process
}

type IncomingNewCharacter struct {
	Name    string
	Map     string
	LenLook uint16
	Look    CharacterLook
}

func (i *IncomingNewCharacter) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	temporaryByte := []byte{}

	handler := func(netes network.Netes) {
		netes.ReadString(&i.Name)
		netes.ReadString(&i.Map)
		netes.ReadUint16(&i.LenLook)
		netes.ReadUint8(&i.Look.SynType)
		netes.ReadUint16(&i.Look.Race)
		netes.ReadUint8(&i.Look.BoatCheck)

		for _, item := range i.Look.Items {
			netes.ReadUint16(&item.Id)
			netes.ReadBytes(&temporaryByte, 160)

			copy(item.Pass[:], temporaryByte)
		}

		netes.ReadUint16(&i.Look.Hair)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		// Correct encoding with Russian support
		i.Map = string(charmap.CP1251_to_UTF8([]byte(i.Map)))

		return []int{support.OP_SERVER_NEWCHA}
	}

	return handler, process
}

type OutcomingNewCharacter struct {
	ErrorCode uint16
}

func (i *OutcomingNewCharacter) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		netes.WriteUint16(i.ErrorCode)

		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		//i.ErrorCode = support.SecretPasswordIncorrect.Code

		return nil
	}

	return handler, process
}
