package main

import (
	"io"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func (s *server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was invoked with a streaming request\n")

	var max int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if req.GetNumber() > max {
			max = req.GetNumber()
		}

		if err := stream.Send(&pb.MaxResponse{Max: max}); err != nil {
			return err
		}
	}
}
