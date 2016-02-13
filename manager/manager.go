package manager

import "github.com/Nyarum/noterius/database"

const (
	CACHE = iota
	DATABASE
)

type Manager struct {
	*database.Database
}

func NewManager(database *database.Database) *Manager {
	return &Manager{Database: database}
}

func (m *Manager) Player(storage int) ManagerPlayer {
	switch storage {
	case CACHE:
		return NewManagerPlayerCache(m)
	case DATABASE:
		return NewManagerPlayerDatabase(m)
	}

	return nil
}
