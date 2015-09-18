package main

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/land"

	"flag"
	"log"
)

func main() {
	configPathFlag := flag.String("config", "resources/config.yml", "A config file for start server")
	flag.Parse()

	app := land.Application{}
	app.Error = core.NewError()
	defer app.Error.GlobalHandler()

	log.Println("Loading config..")
	if err := core.LoadConfig(&app.Config, *configPathFlag); err != nil {
		log.Fatalln("Config is not load, err - ", err)
	}

	log.Println("Loading database..")
	if err := core.LoadDatabase(&app.Database, app.Config.Database.Path); err != nil {
		log.Fatalln("Database is not load, err - ", err)
	}

	log.Printf("Server starting on %v address\n", app.Config.Base.IP+":"+app.Config.Base.Port)
	if err := app.Run(); err != nil {
		log.Fatalln("Server is not started, err - ", err)
	}
}
