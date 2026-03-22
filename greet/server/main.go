package main

import (
	"log"
	"net"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
	"google.golang.org/grpc"
)

var address = "0.0.0.0:50051"

type server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server is listening on %s", address)

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
