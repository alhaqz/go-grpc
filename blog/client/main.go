package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/alhaqz/grpc-learn/blog/proto"
)

var addr string = "localhost:50051"

func main() {	

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect on port: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)
	
	id := createBlog(c)
	readBlog(c, id)
	// readBlog(c, "Invalid Id")
	updateBlog(c, id)
	// updateBlog(c, "Invalid Id")
	listBlog(c)
	deleteBlog(c, id)
}