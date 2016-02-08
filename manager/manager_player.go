package manager

import "github.com/Nyarum/noterius/entitie"

type ManagerPlayer struct{}

func NewManagerPlayer() *ManagerPlayer {
	return &ManagerPlayer
}

func (m *ManagerPlayer) Save(player entitie.Player) {

}
