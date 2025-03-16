package main

import (
	"log"

	"github.com/vague2k/laugh/config"
)

func main() {
	conf, err := config.LoadConfig("")
	if err != nil {
		log.Fatal(err)
	}

	app := NewApp(conf)
	app.Run()

	if err := app.Err(); err != nil {
		log.Fatal(err)
		// app.ExitCLI()
	}
}
