package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc/codes"

	squareroot "../squarerootpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Squareroot(ctx context.Context, req *squareroot.SquarerootRequest) (*squareroot.SquarerootResponse, error) {
	fmt.Printf("Squareroot is invoked")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received a negative number :%v", number),
		)
	}
	return &squareroot.SquarerootResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Squareroot server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	squareroot.RegisterSquarerootServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
