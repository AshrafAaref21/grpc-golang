package main

import (
	"io"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func (s *server) Avg(stream pb.CalculatorService_AvgServer) error {
	var sum int32
	var count int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&pb.AvgResponse{Average: average})
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}
}
