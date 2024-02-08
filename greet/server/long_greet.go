package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func (s *GreetServer) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked with:",)

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF{
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while writing stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
	
}