package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
)

var addr string = "localhost:50051"

func main() {	
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error when loading certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect on port: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	// a := pb.NewAverageServiceClient(conn)
	// d := pb.NewCalculatorServiceClient(conn)
	// e := pb.NewMaxServiceClient(conn)
	// f := pb.NewSqrtServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doSum(d)
	// doCalculateManyTimes(d)
	// doLongGreet(c)
	// doAverage(a)
	// doGreetEveryone(c)
	// doMax(e)
	// doSqrt(f, 10)
	// doSqrt(f, -10)
	// doSqrt(f, 100)
	// doGreetWithDeadline(c, 4* time.Second)
	// doGreetWithDeadline(c, 3* time.Second)
	// doGreetWithDeadline(c, 2* time.Second)
}