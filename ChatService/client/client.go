package main

import (
	"ChatService/api"
	"bufio"
	"flag"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var name string
var client api.PeerClient
var ctx context.Context

func Join() {
	stream, _ := client.Join(context.Background(), &api.JoinMessage{User: name})

	for {
		response, err := stream.Recv()

		if err != nil {
			break
		}

		if response.User == "" {
			log.Default().Printf("Server >> %s", response.Content)
		} else {
			log.Default().Printf("%s >> %s", response.User, response.Content)
		}
	}
}

func Handle(message string) {
	if message == "requestCS" {
		resp, _ := client.Retrieve(ctx, &api.RetrieveMessage{User: name})

		if resp.Message != "" {
			log.Default().Printf("Server >> %s", resp.Message)
		}
	} else if message == "releaseCS" {
		resp, _ := client.Release(ctx, &api.ReleaseMessage{User: name})

		if resp.Message != "" {
			log.Default().Printf("Server >> %s", resp.Message)
		}
	}
}

func main() {
	// Handle flags
	nameFlag := flag.String("name", "", "")

	flag.Parse()
	name = *nameFlag

	// Handle connection
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect! %s", err)
		return
	}

	defer conn.Close()

	client = api.NewPeerClient(conn)
	ctx = context.Background()

	go Join()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go Handle(scanner.Text())
	}
}
