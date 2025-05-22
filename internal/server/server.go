package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Grpc *grpc.Server
}

func New() *Server {
	return &Server{Grpc: grpc.NewServer()}
}

func (s *Server) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to open socket: %w", err)
	}

	log.Printf("Starting gRPC server on %s", addr)
	return s.Grpc.Serve(lis)
}

func (s *Server) Stop() {
	s.Grpc.GracefulStop()
}
