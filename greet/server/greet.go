package main

import (
	"context"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func (s *GreetServer) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoke with %v\n", in);

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}