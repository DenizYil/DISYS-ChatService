package chat

import (
	"golang.org/x/net/context"
	"log"
	"strconv"
)



type Server struct {
	UnimplementedChatServiceServer
}

var clients []ChatService_JoinServer = make([]ChatService_JoinServer, 0)
var lamport int32 = 0

func (s *Server) Broadcast(ctx context.Context, message *Message) (*Empty, error) {

	lamport = lamport + 1
	message.Lamport = lamport

	for _, client := range clients {
		client.Send(message)
	}

	return &Empty{}, nil
}

func (s *Server) Join(message *JoinMessage, stream ChatService_JoinServer) error {

	log.Printf("%s has connected to the chat!", message.User)

	clients = append(clients, stream)

	lamport = lamport + 1
	stream.Send(&Message{User: "", Content: "Welcome " + message.User + "! You have connected to the chat!", Lamport: lamport})

	msg := Message{
		User: "",
		Content: "Participant " + message.User + " joined Chitty-Chat at Lamport time L",
		Lamport: lamport,
	}

	s.Broadcast(nil, &msg)

	for {
		select {
		case <-stream.Context().Done():
			msg := Message{
				User: "",
				Content: "Participant " + message.User + " left Chitty-Chat at Lamport time L",
				Lamport: lamport,
			}
			for i, element := range clients {
				if element == stream {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			s.Broadcast(nil, &msg)
			return nil
		}
	}

}

func (s *Server) Publish(ctx context.Context, message *Message) (*Empty, error) {
	log.Printf("(%s, %s) >> %s", strconv.Itoa(int(message.Lamport)), message.User, message.Content)
	lamport = MaxInt(lamport, message.Lamport) + 1
	return s.Broadcast(ctx, message)
}

func MaxInt(a int32, b int32) int32{
	if a > b {
		return a
	}
	return b
}
