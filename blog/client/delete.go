package main

import (
	"context"
	"log"

	pb "github.com/alhaqz/grpc-learn/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("---updateBlog was invoked---")

	req := &pb.BlogId{Id: id}
	_, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}