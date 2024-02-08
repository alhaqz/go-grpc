package main

import (
	"log"
	"net"

	pb "github.com/alhaqz/grpc-learn/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type GreetServer struct {
	pb.GreetServiceServer
}
type AverageServer struct {
	pb.AverageServiceServer
}

type CalculatorServer struct {
	pb.CalculatorServiceServer
}

type MaxServer struct {
	pb.MaxServiceServer
}
type SqrtServer struct {
	pb.SqrtServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on port: %v\n", err.Error())
	} 
	log.Printf("Listening on port %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewClientTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificate: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))

	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &GreetServer{})
	pb.RegisterCalculatorServiceServer(s, &CalculatorServer{})
	pb.RegisterAverageServiceServer(s, &AverageServer{})
	pb.RegisterMaxServiceServer(s, &MaxServer{})
	pb.RegisterSqrtServiceServer(s, &SqrtServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}