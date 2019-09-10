package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../squareroot"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Squareroot(ctx context.Context, req *squareroot.SquarerootRequest) (*squareroot.SquarerootResponse, error) {
	fmt.Printf("Add is invoked: %v", req)
	a := req.A
	b := req.B
	result := a + b
	res := &addpb.AddResponse{
		Res: result,
	}
	return res, nil
}

func main() {
	fmt.Println("ADD server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	squarerootpb.RegisterSquarerootServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
