package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}
	api.run(api.mount())

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server has failed to start: %v\n", err)
		os.Exit(1)
	}
}
