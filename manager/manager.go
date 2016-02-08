package manager

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Player() *ManagerPlayer {
	return NewManagerPlayer()
}

func (m *Manager) Monster() *ManagerMonster {
	return NewManagerMonster()
}

func (m *Manager) NPC() *ManagerNPC {
	return NewManagerNPC()
}
