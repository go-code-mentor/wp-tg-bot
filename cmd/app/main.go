package main

import (
	"github.com/go-code-mentor/wp-tg-bot/internal/app"
	"github.com/go-code-mentor/wp-tg-bot/internal/config"
	"log"
)

func main() {
	cfg := config.New()
	err := cfg.ParseConfig()
	if err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}

	a := app.New(cfg)
	if err = a.Build(); err != nil {
		log.Fatalf("failed to build app: %s", err)
	}
	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}

	log.Println("app successfully stopped")
}
