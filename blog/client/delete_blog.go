package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, blogID string) {
	log.Printf("Deleting a blog...")
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: blogID})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog was deleted")

}
