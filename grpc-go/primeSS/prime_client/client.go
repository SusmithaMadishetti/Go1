package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"../prime"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Im in Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnt connect:%v", err)
	}
	defer conn.Close()
	c := prime.NewPrimeServiceClient(conn)
	doSS(c)

}
func doSS(c prime.PrimeServiceClient) {
	req := &prime.PrimeRequest{
		Number: 12,
	}
	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling rpc:%v", err)

	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading Stream:%v", err)
		}
		fmt.Println(res.GetResult())
	}

}
