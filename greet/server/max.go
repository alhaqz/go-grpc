package main

import (
	"io"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func (s *MaxServer) Max(stream pb.MaxService_MaxServer) error {
	log.Println("Max was invoked")

	var curMaxNum int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF{
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := req.Number

		if (res > curMaxNum) {
			curMaxNum = res
			err = stream.Send(&pb.MaxResponse{
				Max: res,
			})

			if err != nil {
				log.Fatalf("Error while sending data to clieng: %v\n", err)
			}
		} 
	}
}