package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("Starting Sqrt...")

	req := &pb.SqrtRequest{Number: n}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Sqrt(ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Printf("Error message from server: %s", s.Message())
			log.Printf("Error code from server: %s", s.Code())
		} else {
			log.Printf("Non gRPC error: %v", err)
		}
		log.Fatalf("Error while calling Sqrt RPC: %v", err)
	}
	log.Printf("Response from Sqrt: %v", res.GetResult())

}
