package database

type Character struct {
	Database
}

func NewCharacter(db Database) *Character {
	return &Character{Database: db}
}

func (c *Character) Get() (*Character, error) {
	return c, nil
}

func (c *Character) Save() error {
	return nil
}

func (c *Character) Update() error {
	return nil
}

func (c *Character) Remove() error {
	return nil
}
