package main

import (
	"context"
	"log"
	"math"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Received Sqrt request: %v", req)

	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"%s", "Recieved a negative number: "+string(number),
		)
	}

	result := math.Sqrt(float64(number))
	return &pb.SqrtResponse{Result: result}, nil
}
