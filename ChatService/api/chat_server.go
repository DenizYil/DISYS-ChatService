package api

import (
	"log"
	"sync"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedPeerServer
}

var clients []Peer_JoinServer = make([]Peer_JoinServer, 0)
var queue []Peer_RetrieveServer = make([]Peer_RetrieveServer, 0)
var mutex sync.Mutex
var ctx context.Context

func (s *Server) Broadcast(ctx context.Context, message *Message) (*Empty, error) {

	for _, client := range clients {
		client.Send(message)
	}

	return &Empty{}, nil
}

func (s *Server) Join(message *JoinMessage, stream Peer_JoinServer) error {

	log.Printf("%s has connected to the server!", message.User)

	clients = append(clients, stream)

	stream.Send(&Message{User: "", Content: "Welcome " + message.User + "! You have connected to the network!"})

	msg := Message{
		User:    "",
		Content: "Participant " + message.User + " joined the server",
	}

	s.Broadcast(nil, &msg)

	for {
		select {
		case <-stream.Context().Done():
			msg := Message{
				User:    "",
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
	log.Printf("(%s, %s) >> %s", message.User, message.Content)
	return s.Broadcast(ctx, message)
}

func (s *Server) Retrieve(msg *RetrieveMessage, stream Peer_RetrieveServer) error {

	if len(queue) == 0 {
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been given away"})
	} else {
		s.Broadcast(ctx, &Message{Content: "The critical section is in use!"})
	}

	queue = append(queue, stream)
	mutex.Lock()

	return nil
}

func (s *Server) Release(ctx context.Context, empty *Empty) (*Empty, error) {
	mutex.Unlock()

	queue = RemoveIndex(queue, 0)

	if len(queue) != 0 {
		mutex.Lock()
		s.Broadcast(context.TODO(), &Message{Content: "The Critical Section has been given away as there was someone in the queue waiting for it!"})
	} else {
		s.Broadcast(context.TODO(), &Message{Content: "The Critical Section is now available for anyone to pick up"})
	}

	return &Empty{}, nil
}

func RemoveIndex(s []Peer_RetrieveServer, index int) []Peer_RetrieveServer {
	return append(s[:index], s[index+1:]...)
}
