package main

import (
	"github.com/go-code-mentor/wp-tg-bot/internal/config"
	server "github.com/go-code-mentor/wp-tg-bot/internal/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("failed to pasre config: %s", err)
	}

	srv := server.New()

	go func() {
		if err := srv.Run(cfg.GrpcConnString()); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	srv.Stop()
	log.Println("Server stopped")
}
