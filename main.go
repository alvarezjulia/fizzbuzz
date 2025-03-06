package main

import (
	"log"

	"github.com/alvarezjulia/fizzbuzz/config"
	"github.com/alvarezjulia/fizzbuzz/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	srv := server.NewServer(cfg)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
