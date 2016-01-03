package pill

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
)

type Exit struct {
	err     error
	opcodes []int
}

func (p *Exit) Error() error {
	return p.err
}

func (p *Exit) Opcodes() []int {
	return p.opcodes
}

func (p *Exit) Handler(netes network.Netes) PillFactory {
	return p
}

func (p *Exit) Process(player *entitie.Player) PillFactory {
	player.Connection.Close()

	return p
}
