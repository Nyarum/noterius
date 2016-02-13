package core

import (
	"github.com/Nyarum/migrations"
	log "github.com/Sirupsen/logrus"
)

func Migration1(db migrations.DB) error {
	log.Println("Create user table")

	_, err := db.Exec(`
		DROP TABLE IF EXISTS "public"."user";
		CREATE TABLE "public"."user" (
			"id" int4 NOT NULL,
			"name" varchar(255) COLLATE "default",
			"password" varchar(255) COLLATE "default",
			"created_at" timestamp(6)
		) WITH (OIDS=FALSE);

		ALTER TABLE "public"."user" ADD PRIMARY KEY ("id");
	`)

	return err
}
