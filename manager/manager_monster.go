package manager

import "github.com/Nyarum/noterius/entitie"

type ManagerMonster struct{}

func NewManagerMonster() *ManagerMonster {
	return &ManagerMonster
}

func (m *ManagerMonster) Save(monster entitie.Monster) {

}
