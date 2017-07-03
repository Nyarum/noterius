package core

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Load(dsn string) error {
	db, err := sql.Open("postgres", "postgres://"+dsn)
	if err != nil {
		return err
	}

	_, err = db.Exec("SELECT COUNT(*)")
	if err != nil {
		return err
	}

	d.DB = db

	return nil
}
