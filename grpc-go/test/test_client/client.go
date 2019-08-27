package main

import (
	"fmt"
	"log"

	"../testpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello im client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}
	defer conn.Close()
	c := testpb.NewTestServiceClient(conn)
	fmt.Printf("created client:%f", c)
}
