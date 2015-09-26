package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"fmt"
)

// Database struct
type Database struct {
	DB gorm.DB
}

// NewDatabase method for load database from path
func NewDatabase(database *Database, config *Config) (err error) {
	loadDb, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Database.IP, config.Database.User, config.Database.Password, config.Database.Name))
	defer loadDb.Close()
	if err != nil {
		return
	}

	database.DB = loadDb

	return
}
