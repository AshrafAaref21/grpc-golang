package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func getPrimes(client pb.CalculatorServiceClient, limit int32) {
	stream, err := client.Primes(context.Background(), &pb.PrimesRequest{Limit: limit})
	if err != nil {
		log.Fatalf("Error calling Primes: %v", err)
	}

	log.Printf("Prime numbers up to %d:", limit)
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("%d", resp.GetPrime())
	}
}
