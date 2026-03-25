package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("Starting to do a GreetWithDeadline RPC with timeout: %v", timeout)

	req := &pb.GreetRequest{FirstName: "Ashraf"}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline was exceeded: %v", e.Message())
			} else {
				log.Printf("Unexpected gRPC error: %v", e)
			}
		} else {
			log.Printf("Non gRPC error: %v", err)
		}
		return
	}
	log.Printf("Response from GreetWithDeadline: %v", res.GetMessage())
}
