package main

import (
	"ChatService/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Loading Chitty-Chat service now...")

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
	log.Print("                            Welcome to Chitty-Chat!                            ")
	log.Print("            Users can connect at any time and api with each other!            ")
	log.Print("===============================================================================")

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to server gRPC serevr over port 9000")
	}
}
