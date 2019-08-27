package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../addpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *addpb.AddRequest) (*addpb.AddResponse, error) {
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
	addpb.RegisterAddServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
