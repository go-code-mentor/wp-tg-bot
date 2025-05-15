package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server interface {
	Conn() *grpc.Server
	Run(addr string) error
	Stop()
}

type grpcServer struct {
	grpc *grpc.Server
}

func New() Server {
	return &grpcServer{grpc: grpc.NewServer()}
}

func (s *grpcServer) Conn() *grpc.Server {
	return s.grpc
}

func (s *grpcServer) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on %s", addr)
	return s.grpc.Serve(lis)
}

func (s *grpcServer) Stop() {
	s.grpc.GracefulStop()
}
