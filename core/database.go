package core

import (
	"github.com/Nyarum/migrations"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/pg.v3"
)

// Database struct
type DatabaseInfo struct {
	DB *pg.DB
}

// NewDatabase method for load database from path
func NewDatabaseInfo(config *Config) (database DatabaseInfo, err error) {
	database.DB = pg.Connect(&pg.Options{
		Host:     config.Database.IP,
		User:     config.Database.User,
		Password: config.Database.Password,
		Database: config.Database.Name,
	})

	// Check connect
	_, err = database.DB.Exec(`SELECT SUM(1 + 1)`)
	if err != nil {
		return
	}

	for _, migration := range migrationsWorld {
		err = migrations.Register(migration.Version, migration.Up, migration.Down)
		if err != nil {
			return
		}
	}

	oldVersion, newVersion, err := migrations.Run(database.DB, "up")
	if err != nil {
		return
	}

	if newVersion != oldVersion {
		log.WithFields(log.Fields{
			"old": oldVersion,
			"new": newVersion,
		}).Info("Migrations have applied")
	}

	return
}
