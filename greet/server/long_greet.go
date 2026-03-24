package main

import (
	"io"
	"log"
	"strings"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
)

func (s *server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request...")

	var result strings.Builder
	result.WriteString("Hello ")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error receiving stream from client: %v", err)
			break
		}
		log.Printf("Received request: %v", req)
		result.WriteString(req.GetFirstName() + " ")
	}

	return stream.SendAndClose(&pb.GreetResponse{
		Message: result.String(),
	})
}
