package main

import (
	"io"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func (s *AverageServer) Average(stream pb.AverageService_AverageServer) error {
	log.Printf("Average function was invoked!")

	res := int64(0)
	var number int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF{
			return stream.SendAndClose(&pb.AverageResponse{
				Average: float64(res) / float64(number),
			})
		}

		if err != nil {
			log.Fatalf("Error while writing stream: %v\n", err)
		}

		res += req.Number
		number += 1
	}
}