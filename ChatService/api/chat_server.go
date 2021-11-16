package api

import (
	"golang.org/x/net/context"
	"log"
	"strconv"
	"sync"
)



type Server struct {
	UnimplementedChatServiceServer
}

var clients []ChatService_JoinServer = make([]ChatService_JoinServer, 0)
var queue []ChatService_JoinServer = make([]ChatService_JoinServer, 0)
var mutex sync.Mutex

func (s *Server) Broadcast(ctx context.Context, message *Message) (*Empty, error) {



	for _, client := range clients {
		client.Send(message)
	}

	return &Empty{}, nil
}

func (s *Server) Connect(message *JoinMessage, stream ChatService_JoinServer) error {

	log.Printf("%s has connected to the server!", message.User)

	clients = append(clients, stream)


	stream.Send(&Message{User: "", Content: "Welcome " + message.User + "! You have connected to the network!"})

	msg := Message{
		User: "",
		Content: "Participant " + message.User + " joined the server",
	}

	s.Broadcast(nil, &msg)

	for {
		select {
		case <-stream.Context().Done():
			msg := Message{
				User: "",
				Content: "Participant " + message.User + " left the server",
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
	return s.Broadcast(ctx, message)
}

func (s *Server) Retrieve(ctx context.Context, in *RetrieveMessage, stream ChatService_JoinServer) (*RetrieveReply, error) {


	if len(queue) == 0 {
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been given away"})
		return &RetrieveReply{success: true}
	}else {
		queue = append(queue, stream)
		s.Broadcast(ctx, &Message{Content: "The critical section is in use!"})
	}
	mutex.Lock()
}

func (s *Server) Release(ctx context.Context, in *Empty) (*Empty, error) {
		mutex.Unlock()
}
