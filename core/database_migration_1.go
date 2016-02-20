package core

import (
	"github.com/Nyarum/migrations"
	log "github.com/Sirupsen/logrus"
)

func Migration1(db migrations.DB) error {
	log.Info("Create user table")

	_, err := db.Exec(`
		DROP TABLE IF EXISTS "public"."user";
		CREATE TABLE "public"."user" (
			"id" SERIAL PRIMARY KEY NOT NULL,
			"login" varchar(50) COLLATE "default",
			"password" varchar(32) COLLATE "default",
			"is_active" boolean,
			"created_at" timestamp(6)
		) WITH (OIDS=FALSE);
	`)

	return err
}
