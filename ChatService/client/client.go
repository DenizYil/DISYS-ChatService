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

//	response, _ := client.SendMessage(context.Background(), &api.Message{Content: "Hello from the client!"})
//	log.Printf("Response from server: %s", response.Content)

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

		log.Default().Printf("%s >> %s", response.User, response.Content)
	}
}

func Publish(message string) {

	if message == "requestCS" {
		_, err := client.Retrieve(ctx, &api.RetrieveMessage{User: name})

		if err != nil {
			log.Fatalf("Could not send the message.. Error: %s", err)
		}
	} else if message == "releaseCS" {
		client.Release(ctx, &api.ReleaseMessage{User: name})
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
		log.Fatal("could not connect! %s", err)
		return
	}

	defer conn.Close()

	client = api.NewPeerClient(conn)
	ctx = context.Background()

	go Join()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go Publish(scanner.Text())
	}
}
