package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Rpc *grpc.Server
}

func New() *Server {
	rpc := grpc.NewServer()
	return &Server{Rpc: rpc}
}
func (s *Server) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on %s", addr)
	return s.Rpc.Serve(lis)
}

func (s *Server) Stop() {
	s.Rpc.GracefulStop()
}
