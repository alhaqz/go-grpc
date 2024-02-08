package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Bambang"},
		{FirstName: "Biming"},
		{FirstName: "Bumbung"},
		{FirstName: "Agus"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %s\n",req)
		stream.Send(req)
		time.Sleep(1* time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil{
		log.Fatalf("Error while receiving response from LongGreet:%v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}