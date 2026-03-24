package main

import (
	"io"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func (s *server) GreetEverone(stream pb.GreetService_GreetEveroneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		res := "Hello " + req.GetFirstName() + "!"
		if err := stream.Send(&pb.GreetResponse{Message: res}); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
