package manager

import "github.com/Nyarum/noterius/entitie"

type ManagerPlayer interface {
	Get()
	Save(entitie.Player) error
}
