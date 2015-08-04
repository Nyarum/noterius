package main

import (
	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/land"

	"flag"
	"log"
)

func main() {
	configPathFlag := flag.String("config", "resources/config.yaml", "A config file for start server")
	flag.Parse()

	app := land.Application{}

	if err := core.LoadConfig(&app.Config, *configPathFlag); err != nil {
		log.Fatalln("Config is not load, err - ", err)
	}

	log.Printf("Server starting on %v address\n", app.Config.IP+":"+app.Config.Port)
	if err := app.Run(); err != nil {
		log.Fatalln("Server is not started, err - ", err)
	}
}
