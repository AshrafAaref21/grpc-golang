package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	req := &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		panic(err)
	}
	log.Printf("Sum: %d", res.GetResult())
}
