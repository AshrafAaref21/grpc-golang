package main

import (
	"context"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, blogID string) {
	log.Printf("Updating the blog...")

	blog := &pb.Blog{
		Id:       blogID,
		AuthorId: "Ashraf Aaref",
		Content:  "This is my updated blog post",
		Title:    "My Updated Blog",
	}

	res, err := c.UpdateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog has been updated: %v", res)

}
