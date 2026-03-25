package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readBlog(c pb.BlogServiceClient, blogID string) *pb.Blog {
	log.Printf("Reading a blog...")
	res, err := c.ReadBlog(context.Background(), &pb.BlogId{Id: blogID})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog was read: %v", res)
	log.Printf("Blog Title: %v", res.GetTitle())
	log.Printf("Blog Author: %v", res.GetAuthorId())
	log.Printf("Blog Content: %v", res.GetContent())

	oid, err := primitive.ObjectIDFromHex(res.GetId())
	if err != nil {
		log.Fatalf("Error while converting to OID: %v", err)
	}

	fmt.Printf("The OID is: %v\n", oid)

	return res

}
