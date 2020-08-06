package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server is the struct for server properties
type Server struct {
}

// SayHello is the basic function for rpc calls
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
