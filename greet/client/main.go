package main

import (
	"log"
	"time"

	pb "github.com/AshrafAaref21/grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Printf("Connected to server at %s", addr)

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	log.Printf("Finished unary RPC...")

	doGreetManyTimes(c)
	log.Printf("Finished streaming RPC...")

	doLongGreet(c)
	log.Printf("Finished client streaming RPC...")

	doGreetEveryone(c)
	log.Printf("Finished bidirectional streaming RPC...")

	doGreetWithDeadline(c, 5*time.Second) // should complete successfully
	doGreetWithDeadline(c, 1*time.Second) // should timeout
}
