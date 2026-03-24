package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("Starting to do a LongGreet RPC...")

	req := []*pb.GreetRequest{
		{FirstName: "Ashraf"},
		{FirstName: "Aaref"},
		{FirstName: "Golang"},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, r := range req {
		log.Printf("Sending request: %v", r)
		if err := stream.Send(r); err != nil {
			log.Fatalf("Error while sending request to LongGreet: %v", err)
		}
		time.Sleep(300 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet response: %v", resp.GetMessage())
}
