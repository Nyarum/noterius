package manager

import "github.com/Nyarum/noterius/entitie"

type ManagerNPC struct{}

func NewManagerNPC() *ManagerNPC {
	return &ManagerNPC
}

func (m *ManagerNPC) Save(npc entitie.NPC) {

}
