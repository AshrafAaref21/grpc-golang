package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AshrafAaref21/grpc-golang/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreateBlog(ctx context.Context, req *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked")

	data := BlogItem{
		AuthorID: req.GetAuthorId(),
		Content:  req.GetContent(),
		Title:    req.GetTitle(),
	}

	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Internal Server Error while inserting into db. Error: %v", err),
		)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Cannot convert to OID. Error: %v", err),
		)
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}

func (s *server) ReadBlog(ctx context.Context, req *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v", req.GetId())

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"%s", fmt.Sprintf("Cannot parse ID. Error: %v", err),
		)
	}

	filter := bson.M{"_id": oid}

	var result BlogItem
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"%s", fmt.Sprintf("Cannot find blog with specified ID. Error: %v", err),
		)
	}

	return documentToBlog(&result), nil
}

func (s *server) UpdateBlog(ctx context.Context, req *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v", req.GetId())

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"%s", fmt.Sprintf("Cannot parse ID. Error: %v", err),
		)
	}

	filter := bson.M{"_id": oid}

	data := BlogItem{
		AuthorID: req.GetAuthorId(),
		Content:  req.GetContent(),
		Title:    req.GetTitle(),
	}

	update := bson.M{"$set": data}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Cannot update blog. Error: %v", err),
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *server) ListBlogs(req *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Printf("ListBlogs was invoked")

	cursor, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Error while listing blogs. Error: %v", err),
		)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var blog BlogItem
		if err := cursor.Decode(&blog); err != nil {
			return status.Errorf(
				codes.Internal,
				"%s", fmt.Sprintf("Error while decoding blog. Error: %v", err),
			)
		}
		if err := stream.Send(documentToBlog(&blog)); err != nil {
			return status.Errorf(
				codes.Internal,
				"%s", fmt.Sprintf("Error while sending blog. Error: %v", err),
			)
		}
	}

	return nil
}

func (s *server) DeleteBlog(ctx context.Context, req *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v", req.GetId())

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"%s", fmt.Sprintf("Cannot parse ID. Error: %v", err),
		)
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Cannot delete blog. Error: %v", err),
		)
	}

	return &emptypb.Empty{}, nil
}
