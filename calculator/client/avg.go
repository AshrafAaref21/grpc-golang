package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	stream, err := c.Avg(context.Background())
	if err != nil {
		panic(err)
	}

	numbers := []int32{1, 2, 3, 4, 5}
	for _, number := range numbers {
		if err := stream.Send(&pb.AvgRequest{Number: number}); err != nil {
			panic(err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	log.Printf("Average: %f", resp.Average)
}
