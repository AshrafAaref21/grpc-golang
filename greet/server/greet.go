package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("Greet Function was invoked with %v\n", req)

	return &pb.GreetResponse{
		Message: "Hello " + req.GetFirstName(),
	}, nil
}
