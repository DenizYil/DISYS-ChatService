package main

import (
	"ChatService/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Loading Mutual Exclusion Distribution service now...")

	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		log.Fatalf("TCP failed to listen... %s", err)
		return
	}

	log.Print("Listener registered - setting up server now...")

	s := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterPeerServer(grpcServer, &s)

	log.Print("===============================================================================")
	log.Print(" ")
	log.Print("             Welcome to the Mutual Exclusion Destribution service!             ")
	log.Print(" ")
	log.Print("===============================================================================")

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to server gRPC serevr over port 9000")
	}
}
