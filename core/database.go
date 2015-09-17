package core

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Database struct with pointer to LevelDB
type Database struct {
	DB *sql.DB
}

// LoadDatabase method for load LevelDB database from path
func LoadDatabase(database *Database, path string) (err error) {
	loadDb, err := sql.Open("postgres", "user=nato password=natodefault dbname=noterius sslmode=verify-full")
	defer loadDb.Close()
	if err != nil {
		return
	}

	database.DB = loadDb

	return
}
