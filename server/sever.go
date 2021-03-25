package server

import (
	"context"
	"log"

	"github.com/kalpitpant/understanding-twirp/chat"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.Message{Body: "Hello From the Server!"}, nil
}
