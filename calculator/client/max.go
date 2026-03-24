package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("Starting to do a Max Unary RPC...\n")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Max RPC: %v", err)
	}

	numbers := []int32{1, 5, 3, 6, 2, 20}
	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		if err := stream.Send(&pb.MaxRequest{Number: number}); err != nil {
			log.Fatalf("Error while sending data to server: %v", err)
		}

		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while receiving data from server: %v", err)
		}
		log.Printf("Received max: %v\n", res.GetMax())

		time.Sleep(300 * time.Millisecond)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error while closing the stream: %v", err)
	}
}
