package incoming

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/interface"
	"github.com/Nyarum/noterius/library/network"
)

type ExitCrumb struct{}

func (e *ExitCrumb) PreHandler(netes network.Netes) interfaces.PillDecoder {
	return e
}

func (e *ExitCrumb) Process(player entitie.Player) ([]int, error) {
	player.Connection.Close()

	return nil, nil
}
