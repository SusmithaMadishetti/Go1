package main

import (
	"fmt"
	"log"
	"net"

	"../prime"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Prime(req *prime.PrimeRequest, stream prime.PrimeService_PrimeServer) error {
	fmt.Printf("prime is invoked : %v", req)
	number := req.GetNumber()
	k := int64(2)
	for number > 1 {
		if number%k == 0 {
			stream.Send(&prime.PrimeResponse{
				Result: k,
			})
			number = number / k
		} else {
			k++
			//fmt.Printf("divisior increased :%v", k)
		}
	}
	return nil
}

func main() {
	fmt.Println("im in server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("couldn't listen to server :%v", err)
	}
	s := grpc.NewServer()
	prime.RegisterPrimeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%v", err)
	}
}
