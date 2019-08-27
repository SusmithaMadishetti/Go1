package main

import (
	"fmt"
	"log"
	"net"

	"../testpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("hello world")
	lis, err := net.Listen("tcp", "0.0.0.0:50051") //50051 is default for grpc
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
