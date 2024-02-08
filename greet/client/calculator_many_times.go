package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func doCalculateManyTimes(d pb.CalculatorServiceClient) {
	log.Println("doCalculateManyTimes was invoked!")

	req := &pb.CalculatorRequest{
		FirstInteger: 1200,
		SecondInteger: 1200,
	}

	stream, err := d.CalculatorManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling CalculateManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("CalculateManyTimes: %s\n", msg)
	}
}