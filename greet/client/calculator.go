package main

import (
	"context"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func doSum(d pb.CalculatorServiceClient) {
	log.Println("doSum was invoked!")
	res, err := d.Calculator(context.Background(), &pb.CalculatorRequest{
		FirstInteger: 10,
		SecondInteger: 3,
	})

	if err != nil  {
		log.Fatalf("Could not sum: %v", err)
	}

	log.Printf("Summing: %v", res.Sum)
 }