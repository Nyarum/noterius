package interfaces

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
)

type PillEncoder interface {
	Process() PillEncoder
	PostHandler(network.Netes) string
}

type PillDecoder interface {
	PreHandler(network.Netes) PillDecoder
	Process(entitie.Player) ([]int, error)
}
