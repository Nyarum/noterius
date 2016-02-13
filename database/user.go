package database

import "time"

type User struct {
	ID        int64
	Login     string
	Password  string
	CreatedAt time.Time

	Database
}

func NewUser(db Database) *User {
	return &User{Database: db}
}

func (u *User) Get(ID int64) (*User, error) {
	_, err := u.Database.DB.QueryOne(u, `SELECT * FROM "public"."user" WHERE id = ?`, ID)

	return u, err
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
