package main

import (
	"log"

	env "github.com/mohitsolanki026/social/internal"
)

func main() {
	app := &application{
		config: config{
			addr: env.GetEnv("ADDR",":8000"),
		},
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}