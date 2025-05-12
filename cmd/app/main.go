package main

import (
	server "github.com/go-code-mentor/wp-tg-bot/internal/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := server.New()

	go func() {
		if err := srv.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	srv.Stop()
	log.Println("Server stopped")
}
