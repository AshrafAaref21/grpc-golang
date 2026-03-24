package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("Starting to do a GreetEveryone RPC...")

	stream, err := c.GreetEverone(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Ashraf"},
		{FirstName: "John"},
		{FirstName: "Jane"},
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		log.Printf("Response from GreetEveryone: %s", res.GetMessage())
	}
}
