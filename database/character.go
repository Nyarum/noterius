package database

import (
	"time"
)

type Character struct {
	Id        int64
	Name      string
	Level     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCharacter() *Character {
	return &Character{}
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
