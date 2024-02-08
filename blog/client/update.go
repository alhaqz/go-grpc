package main

import (
	"context"
	"log"

	pb "github.com/alhaqz/grpc-learn/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Not Elham",
		Title: "A New Title",
		Content: "A New Content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}