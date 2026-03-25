package main

import (
	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func documentToBlog(document *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       document.ID.Hex(),
		AuthorId: document.AuthorID,
		Content:  document.Content,
		Title:    document.Title,
	}
}
