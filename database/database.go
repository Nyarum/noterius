package database

import "github.com/Nyarum/noterius/core"

type Database struct {
	*core.DatabaseInfo
}

func NewDatabase(databaseInfo *core.DatabaseInfo) *Database {
	return &Database{DatabaseInfo: databaseInfo}
}

func (d *Database) Character() *Character {
	return NewCharacter(*d)
}

func (d *Database) User() *User {
	return NewUser(*d)
}
