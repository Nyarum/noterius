package core

import "github.com/Nyarum/migrations"

var migrationsWorld []migrations.Migration = []migrations.Migration{
	{1, Migration1, NilDown},
}

func NilDown(db migrations.DB) error {
	return nil
}
