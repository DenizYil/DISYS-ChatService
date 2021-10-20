package chat

import (
	"golang.org/x/net/context"
	"log"
	"sync"
)

type Server struct {
	UnimplementedChatServiceServer
}

var clients []ChatService_JoinServer = make([]ChatService_JoinServer, 0)

func (s *Server) Broadcast(ctx context.Context, message *Message) (*Empty, error) {

	for _, client := range clients {
		client.Send(message)
	}

	return &Empty{}, nil
}

func (s *Server) Join(message *JoinMessage, stream ChatService_JoinServer) error {

	log.Printf("%s has connected to the chat!", message.User)

	clients = append(clients, stream)

	stream.Send(&Message{User: "", Content: "Welcome " + message.User + "! You have connected to the chat!"})

	lock := sync.Mutex{}
	lock.Lock()
	lock.Lock()

	return nil
}


func (s *Server) Publish(ctx context.Context, message *Message) (*Empty, error) {
	log.Printf("%s >> %s", message.User, message.Content)
	return s.Broadcast(ctx, message)
}
