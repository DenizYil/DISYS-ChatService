package api

import (
	"golang.org/x/net/context"
	"log"
	"sync"
)



type Server struct {
	UnimplementedPeerServer
}

var clients []Peer_JoinServer = make([]Peer_JoinServer, 0)
var queue []Peer_ReleaseServer = make([]Peer_ReleaseServer, 0)
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
	log.Printf("(%s, %s) >> %s", message.User, message.Content)
	return s.Broadcast(ctx, message)
}

func (s *Server) Retrieve(msg *RetrieveMessage, stream Peer_RetrieveServer) error{

	mutex.Lock()
	if len(queue) == 0 {
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been given away"})
		return nil
	}else {
		queue = append(queue, stream)
		s.Broadcast(ctx, &Message{Content: "The critical section is in use!"})
		return nil
	}


}

func (s *Server) Release(empty *Empty) error {
		mutex.Unlock()
		RemoveIndex(queue, stream)
		return nil
}

func RemoveIndex(queue []Peer_ReleaseServer, stream Peer_ReleaseServer) []Peer_ReleaseServer {
	for i := 0; i < len(queue); i++ {
		if queue[i] == stream{
			return append(queue[:i], queue[i+1:]...)
		}
	}
	return queue
}
