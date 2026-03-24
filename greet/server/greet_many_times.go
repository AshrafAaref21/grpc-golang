package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func (s *server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)
	for i := range 10 {
		res := fmt.Sprintf("Hello %s, number %d", in.GetFirstName(), i)
		if err := stream.Send(&pb.GreetResponse{Message: res}); err != nil {
			return err
		}
		time.Sleep(300 * time.Millisecond)
	}
	return nil
}
