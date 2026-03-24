package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func (s *server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum Function was invoked with %v\n", req)
	result := req.GetFirstNumber() + req.GetSecondNumber()
	return &pb.SumResponse{Result: result}, nil
}
