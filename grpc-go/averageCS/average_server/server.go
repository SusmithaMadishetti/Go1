package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"../averagepb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("I'm in Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen :%v", err)
	}
	s := grpc.NewServer()
	averagepb.RegisterAverageServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
func (*server) Average(stream averagepb.AverageService_AverageServer) error {
	fmt.Printf("average function is invoked:")
	sum := int64(0)
	count := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&averagepb.AverageResponse{
				Result: sum / count,
			})

		}
		if err != nil {
			log.Fatalf("error while reading client stream:%v", err)
		}
		sum += req.GetNumber()
		count++
	}
}
