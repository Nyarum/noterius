package database

type User struct {
	Database
}

func NewUser(db Database) *User {
	return &User{Database: db}
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
