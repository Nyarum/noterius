package packet

import (
	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
)

type IncomingKitbagtempSync struct {
}

func (i *IncomingKitbagtempSync) Packet(db *database.Database) (func(netes network.Netes), func(player *entitie.Player) []int) {
	handler := func(netes network.Netes) {
		return
	}

	process := func(player *entitie.Player) []int {
		if player.Error != nil {
			// Here was gopher...
		}

		return nil
	}

	return handler, process
}
