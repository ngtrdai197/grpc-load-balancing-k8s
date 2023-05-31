package grpc_server

import (
	"context"
	"fmt"
	"log"

	"github.com/ngtrdai197/grpc-lb/pkg/pb"
)

func (server *Server) SendMessage(
	ctx context.Context,
	req *pb.Message,
) (*pb.Message, error) {

	log.Printf("request msg \"%s\"", req.GetMsg())

	r := &pb.Message{
		Msg: fmt.Sprintf("received msg"),
	}
	return r, nil
}
