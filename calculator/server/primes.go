package main

import (
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
)

func (*server) Primes(req *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Received Primes RPC: %v\n", req)

	n := req.GetLimit()
	divisor := int32(2)

	for n > 1 {
		if n%divisor == 0 {
			if err := stream.Send(&pb.PrimesResponse{Prime: divisor}); err != nil {
				return err
			}
			n /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
