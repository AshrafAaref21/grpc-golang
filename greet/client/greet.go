package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("Invoking Greet RPC...")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Ashraf",
	})
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %s", res.GetMessage())
}
