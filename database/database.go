package database

type Database struct {
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Character() DatabaseFactory {
	return NewCharacter(*d)
}

func (d *Database) User() DatabaseFactory {
	return NewUser(*d)
}
