package api

import (
	"log"
	"strconv"
	"sync"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedPeerServer
}

var clients []Peer_JoinServer = make([]Peer_JoinServer, 0)
var queue []Peer_RetrieveServer = make([]Peer_RetrieveServer, 0)
var queueNames []string = make([]string, 0)
var mutex sync.Mutex
var ctx context.Context
var currentHolder string

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

	for i := 0; i < len(queueNames); i++ {
		if msg.User == queueNames[i] || currentHolder == msg.User {
			 s.Broadcast(ctx, &Message{User:msg.User, Content: " You are already in the queue " + strconv.Itoa(len(queueNames))})
			return nil
		}
	}

	if len(queue) == 0 {
		currentHolder = msg.User
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been given away to: " + msg.User})
	} else {
		queueNames = append(queueNames, msg.User)
		s.Broadcast(ctx, &Message{User:msg.User, Content: "The critical section is in use! You have been added to the queue and are waiting: Current number - " + strconv.Itoa(len(queueNames))})
	}

	queue = append(queue, stream)
	mutex.Lock()

	return nil
}

func (s *Server) Release(ctx context.Context, msg *ReleaseMessage) (*Empty, error) {

	if queueNames[0] != currentHolder || !isinQue(msg) {
		s.Broadcast(context.TODO(), &Message{User: msg.User, Content: "You just tried to do something illegal! You cannot release until you entered the Critical Section!"})
		return &Empty{}, nil
	}

	mutex.Unlock()

	if len(queue) != 0 {
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been released by: " + queueNames[0]})
		queue = RemovePeer(queue, 0)
		RemovePeerName(queueNames, 0)
		s.Broadcast(context.TODO(), &Message{Content: "The Critical Section is now available for the next in the queue to pick up!"})
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been given away to: " + queueNames[0]})
		//s.Broadcast(context.TODO(), &Message{Content: "The Critical Section has been given away as there was someone in the queue waiting for it!"})
	} else {
		s.Broadcast(context.TODO(), &Message{Content: "The Critical Section is now available for anyone to pick up"})
	}

	return &Empty{}, nil
}

func RemovePeer(s []Peer_RetrieveServer, index int) []Peer_RetrieveServer {
	return append(s[:index], s[index+1:]...)
}
func RemovePeerName(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func isinQue(user *ReleaseMessage) bool{
	for i := 0; i < len(queueNames); i++ {
		if user.User == queueNames[i] {
			return true
		}
	}
	return false
}
