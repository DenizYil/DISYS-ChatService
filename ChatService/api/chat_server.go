package api

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedPeerServer
}

var clients []Peer_JoinServer = make([]Peer_JoinServer, 0)

var queue []string = make([]string, 0)

var mutex sync.Mutex
var ctx context.Context = context.TODO()

var currentHolder string = ""

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

	s.Broadcast(ctx, &msg)

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
			s.Broadcast(ctx, &msg)
			return nil
		}
	}
}

func (s *Server) Retrieve(ctx context.Context, msg *RetrieveMessage) (*ResponseMessage, error) {
	if currentHolder == msg.User {
		return &ResponseMessage{Message: "You already have the key!"}, nil
	}

	for i, name := range queue {
		if msg.User == name {
			return &ResponseMessage{Message: fmt.Sprintf("You are already in the queue... There are %s/%s people in front of you.", strconv.Itoa(len(queue)-(i+1)), strconv.Itoa(len(queue)))}, nil
		}
	}

	if currentHolder == "" {
		currentHolder = msg.User
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been given away to: " + msg.User})
	} else {
		queue = append(queue, msg.User)
		s.Broadcast(ctx, &Message{Content: fmt.Sprintf("%s has been added to the queue as the Critical Section is currently occupied by %s", msg.User, currentHolder)})
	}

	mutex.Lock()
	return &ResponseMessage{Message: ""}, nil
}

func (s *Server) Release(ctx context.Context, msg *ReleaseMessage) (*ResponseMessage, error) {
	if currentHolder != msg.User {
		return &ResponseMessage{Message: "You don't have the critical section"}, nil
	}

	mutex.Unlock()

	s.Broadcast(ctx, &Message{Content: "The Critical Section has been released by: " + msg.User})

	if len(queue) != 0 {
		s.Broadcast(ctx, &Message{Content: "The Critical Section is now available for the next in the queue to pick up!"})
		s.Broadcast(ctx, &Message{Content: "The Critical Section has been given away to: " + queue[0]})

		currentHolder = queue[0]

		queue = RemovePeerName(queue, 0)
	} else {
		s.Broadcast(ctx, &Message{Content: "The Critical Section is now available for anyone to pick up"})
		currentHolder = ""
	}

	return &ResponseMessage{Message: ""}, nil
}

func RemovePeerName(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
