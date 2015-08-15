package core

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Database struct {
	DB *leveldb.DB
}

func LoadDatabase(database *Database, path string) (err error) {
	loadDb, err := leveldb.OpenFile(path, nil)
	defer loadDb.Close()

	if err != nil {
		return
	}

	database.DB = loadDb

	return
}
