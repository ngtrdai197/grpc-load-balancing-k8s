package grpc_server

import (
	"log"
	"net"

	"github.com/ngtrdai197/grpc-lb/config"
	"github.com/ngtrdai197/grpc-lb/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedMessageServiceServer
	config *config.Config
}

// NewServer creates a new gRPC server.
func NewServer(c *config.Config) *Server {
	server := &Server{
		config: c,
	}

	return server
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen, error=%v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server, error=%v", err)
	}
}
