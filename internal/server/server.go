package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server interface {
	Run(addr string) error
	Stop()
}

type Grpc struct {
	Server *grpc.Server
}

func New() *Grpc {
	return &Grpc{Server: grpc.NewServer()}
}

func (s *Grpc) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on %s", addr)
	return s.Server.Serve(lis)
}

func (s *Grpc) Stop() {
	s.Server.GracefulStop()
}
