package main

import (
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/calculator/proto"
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

	c := pb.NewCalculatorServiceClient(conn)

	doSum(c)
	log.Println("------------------------------------------------")

	getPrimes(c, 1581578712)
	log.Println("------------------------------------------------")

	doAvg(c)
	log.Println("------------------------------------------------")

	doMax(c)
	log.Println("------------------------------------------------")

	doSqrt(c, 16)
	doSqrt(c, -16)

}
