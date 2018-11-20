package main

import (
	"log"

	myapp "github.com/vlam1/acceptance_test_example/app"
	"github.com/vlam1/acceptance_test_example/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	srv := myapp.NewServer(cfg)

	srv.Start()
}
