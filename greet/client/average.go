package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func doAverage(a pb.AverageServiceClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := a.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doAverage: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %s\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from doAverage: %v\n", err)
	}

	log.Printf("Average: %v\n", res.Average)
}