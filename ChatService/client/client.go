package main

import (
	"ChatService/chat"
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"sync/atomic"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//	response, _ := client.SendMessage(context.Background(), &chat.Message{Content: "Hello from the client!"})
//	log.Printf("Response from server: %s", response.Content)

var name string
var client chat.ChatServiceClient
var ctx context.Context
var lamport int32 = 0

func Join() {
	stream, _ := client.Join(context.Background(), &chat.JoinMessage{User: name})

	for {
		response, err := stream.Recv()

		if err != nil {
			break
		}

		atomic.StoreInt32(&lamport, MaxInt(lamport, response.Lamport)+1)

		if response.User == "" {
			log.Default().Printf("(%s) >> %s", strconv.Itoa(int(lamport)), response.Content)
			continue
		}

		log.Default().Printf("(%s, %s) >> %s", strconv.Itoa(int(lamport)), response.User, response.Content)
	}
}

func MaxInt(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func Publish(message string) {
	if len(message) > 128 {
		log.Print("The message must not be above 128 characters!")
		return
	}

	if len(message) == 0 {
		log.Print("The message cannot be empty!")
		return
	}

	atomic.AddInt32(&lamport, 1)
	_, err := client.Publish(ctx, &chat.Message{User: name, Content: message, Lamport: lamport})

	if err != nil {
		log.Fatalf("Could not send the message.. Error: %s", err)
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

	client = chat.NewChatServiceClient(conn)
	ctx = context.Background()

	go Join()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go Publish(scanner.Text())
	}
}
