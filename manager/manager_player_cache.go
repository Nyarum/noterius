package manager

import (
	"github.com/Nyarum/noterius/entitie"
	log "github.com/Sirupsen/logrus"
)

type ManagerPlayerCache struct {
	*Manager
}

func NewManagerPlayerCache(manager *Manager) *ManagerPlayerCache {
	return &ManagerPlayerCache{Manager: manager}
}

func (m *ManagerPlayerCache) Get() {

}

func (m *ManagerPlayerCache) Save(player entitie.Player) error {
	log.Info("Test from cache in player manager")

	return nil
}
