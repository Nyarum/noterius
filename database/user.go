package database

import "time"

type User struct {
	ID        int64
	Login     string
	Password  string
	IsActive  bool
	CreatedAt time.Time

	Database
}

func NewUser(db Database) *User {
	return &User{Database: db}
}

func (u *User) GetByID(id int64) (*User, error) {
	_, err := u.Database.DB.QueryOne(u, `SELECT * FROM "user" WHERE id = ?;`, id)

	return u, err
}

func (u *User) GetByName(login string) (*User, error) {
	_, err := u.Database.DB.QueryOne(u, `SELECT * FROM "user" WHERE login = ?;`, login)

	return u, err
}

func (u *User) Save() error {
	_, err := u.Database.DB.QueryOne(u, `
		INSERT INTO "user" ("login", "password", "is_active", "created_at") 
		VALUES (?login, ?password, ?is_active, current_timestamp)
		RETURNING id;
	`, u)

	return err
}

func (u *User) Update(id int64) error {
	u.ID = id

	_, err := u.Database.DB.QueryOne(u, `
		UPDATE "user" SET is_active = ?is_active 
		WHERE id = ?id
		RETURNING id;
	`, u)

	return err
}

func (u *User) Remove(id int64) error {
	_, err := u.Database.DB.Exec(`DELETE FROM "user" WHERE id = ?;`, id)

	return err
}
