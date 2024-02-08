package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Elham",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Errpr while calling GreatManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg)
	}
}