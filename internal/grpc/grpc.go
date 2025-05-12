package grpc

import (
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func New() *Server {
	grpcServer := grpc.NewServer()

	pingService := service.NewPingService()
	api.RegisterPingerServer(grpcServer, pingService)

	return &Server{
		grpcServer: grpcServer,
	}
}

func (s *Server) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on %s", addr)
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
