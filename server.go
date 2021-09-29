package main

import (
	"context"
	"fmt"
	"log"
	"net"

	chat "github.com/sdinescu-dn/grpc_tests/proto"
	"google.golang.org/grpc"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.Message{Body: "Hello From the Server!"}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
