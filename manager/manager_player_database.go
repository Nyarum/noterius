package manager

import "github.com/Nyarum/noterius/entitie"

type ManagerPlayerDatabase struct {
	*Manager
}

func NewManagerPlayerDatabase(manager *Manager) *ManagerPlayerDatabase {
	return &ManagerPlayerDatabase{Manager: manager}
}

func (m *ManagerPlayerDatabase) Get() {

}

func (m *ManagerPlayerDatabase) Save(player entitie.Player) error {
	return nil
}
