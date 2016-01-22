package database

import (
	"time"
)

type User struct {
	Id          int64
	Login       string
	Password    string
	CharacterId int64
	Characters  []Character
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser() *User {
	return &User{}
}

func (u *User) Get() (*User, error) {
	return u, nil
}

func (u *User) Save() error {
	return nil
}

func (u *User) Update() error {
	return nil
}

func (u *User) Remove() error {
	return nil
}
