package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Printf("Creating a new blog...")

	blog := &pb.Blog{
		AuthorId: "Ashraf Aaref",
		Content:  "This is my first blog post",
		Title:    "My First Blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog has been created: %v", res.GetId())

	oid, err := primitive.ObjectIDFromHex(res.GetId())
	if err != nil {
		log.Fatalf("Error while converting to OID: %v", err)
	}

	fmt.Printf("The OID is: %v\n", oid)
	return res.GetId()
}
