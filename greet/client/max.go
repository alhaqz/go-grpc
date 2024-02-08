package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func doMax(e pb.MaxServiceClient) {
	log.Println("doMax was invoked!")

	stream, err := e.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest {
		{Number: 1},
		{Number: 2},
		{Number: 1},
		{Number: 3},
		{Number: 2},
		{Number: 5},
		{Number: 2},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF{
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			
			log.Printf("Received: %v\n", res.Max)
		}
		close(waitc)
	} ()
	<- waitc
}