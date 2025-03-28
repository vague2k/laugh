package main

import (
	"log"

	"github.com/vague2k/laugh/config"
	"github.com/vague2k/laugh/database"
)

func main() {
	db, err := database.New("")
	if err != nil {
		log.Fatal(err)
	}

	conf, err := config.LoadConfig("")
	if err != nil {
		log.Fatal(err)
	}

	app := NewApp(conf, db)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
