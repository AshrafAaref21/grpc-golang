package main

import (
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
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

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	// readBlog(c, "wrong-id")
	updateBlog(c, id)
	listBlogs(c)
	deleteBlog(c, id)
}
