package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kalpitpant/understanding-twirp/chat"
	"github.com/kalpitpant/understanding-twirp/server"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
