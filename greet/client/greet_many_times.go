package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("Starting to do a streaming RPC...")

	req := &pb.GreetRequest{
		FirstName: "Ashraf",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error while receiving from GreetManyTimes stream: %v", err)
			break
		}
		log.Printf("Response from GreetManyTimes: %s", msg.GetMessage())
	}
}
