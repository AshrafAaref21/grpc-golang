package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Printf("Listing all blogs...")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs RPC: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err != nil {
			if err != io.EOF {
				log.Printf("Error while receiving stream: %v", err)
			}
			break
		}

		log.Printf("Blog: %v", res)

		oid, err := primitive.ObjectIDFromHex(res.GetId())
		if err != nil {
			log.Fatalf("Error while converting to OID: %v", err)
		}

		fmt.Printf("The OID is: %v\n", oid)
	}
}
