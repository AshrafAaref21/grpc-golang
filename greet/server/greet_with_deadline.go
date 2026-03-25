package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received GreetWithDeadline request: %v", req)

	for range 3 {
		if ctx.Err() == context.Canceled {
			log.Println("Client canceled the request")
			return nil, status.Error(codes.Canceled, "Client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetFirstName()
	result := "Hello " + firstName
	return &pb.GreetResponse{Message: result}, nil
}
