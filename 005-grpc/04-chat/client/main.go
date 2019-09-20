package main

import (
	"context"
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/GoesToEleven/golang-project/005-grpc/04-chat/chat"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Must have a url to connect to as the first argument, and a username as the second argument")
		return
	}

	ctx := context.Background()

	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := chat.NewChatClient(conn)
	stream, err := c.Chat(ctx)
	if err != nil {
		panic(err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			} else if err != nil {
				panic(err)
			}
			fmt.Println(msg.User + ": " + msg.Message)
		}
	}()

	fmt.Println("Connection established, type \"quit\" or use ctrl+c to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "quit" {
			err := stream.CloseSend()
			if err != nil {
				panic(err)
			}
			break
		}

		err := stream.Send(&chat.ChatMessage{
			User: os.Args[2],
			Message: msg,
		})
		if err != nil {
			panic(err)
		}
	}

	<-waitc
}
