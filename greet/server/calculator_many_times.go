package main

import (
	"fmt"
	"log"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

func (s *CalculatorServer) CalculatorManyTimes(in *pb.CalculatorRequest, stream pb.CalculatorService_CalculatorManyTimesServer) error {
	log.Printf("CalculatorManyTImes function was invoked with: %v\n", in)

	var k int64 = 2;
	N := in.FirstInteger + in.SecondInteger;

	log.Printf("Calculate number: %v\n", N)
	var factors []int64 

	for N > 1 {
		if N%k == 0 {
			fmt.Println(k)
			N = N / k    
			factors = append(factors, k)

			stream.Send(&pb.CalculatorResponse{
				Sum: k,
			})
			} else {
				k = k + 1
		}
	}

	fmt.Printf("Total factor is: %v", factors)
	return nil
}