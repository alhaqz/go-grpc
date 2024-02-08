package main

import (
	"context"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func (s *CalculatorServer) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoke with %v\n", in)

	return &pb.CalculatorResponse{
		Sum: in.FirstInteger + in.SecondInteger,
	}, nil
}