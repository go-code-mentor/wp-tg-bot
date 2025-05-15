package main

import (
	"github.com/go-code-mentor/wp-tg-bot/internal/app"
	"github.com/go-code-mentor/wp-tg-bot/internal/config"
	"github.com/go-code-mentor/wp-tg-bot/internal/server"
	"log"
)

func main() {
	cfg := config.New()
	err := cfg.ParseConfig()
	if err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}

	srv := server.New()

	a := app.New(cfg, srv)
	a.Build()
	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}

	log.Println("app successfully stopped")
}
