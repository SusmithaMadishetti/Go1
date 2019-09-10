package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"../maxpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen :%v", err)
	}
	s := grpc.NewServer()
	maxpb.RegisterMaxServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func (*server) Max(stream maxpb.MaxService_MaxServer) error {
	fmt.Printf("Max function is invoked")
	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client:%v", err)
			return err
		}
		num := req.GetNumber()
		if num > maximum {
			maximum = num
			sendErr := stream.Send(&maxpb.MaxResponse{
				Result: maximum,
			})
			if sendErr != nil {
				log.Fatalf("error while sending data to client:%v", sendErr)
				return sendErr
			}
		}

	}
}
