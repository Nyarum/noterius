package core

import (
	"github.com/syndtr/goleveldb/leveldb"
)

// Database struct with pointer to LevelDB
type Database struct {
	DB *leveldb.DB
}

// LoadDatabase method for load LevelDB database from path
func LoadDatabase(database *Database, path string) (err error) {
	loadDb, err := leveldb.OpenFile(path, nil)
	defer loadDb.Close()

	if err != nil {
		return
	}

	database.DB = loadDb

	return
}
