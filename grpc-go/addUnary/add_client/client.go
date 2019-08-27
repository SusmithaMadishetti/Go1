package main

import (
	"context"
	"fmt"
	"log"

	"../addpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello,I'm client!!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect:%v", err)
	}
	defer cc.Close()
	c := addpb.NewAddServiceClient(cc)
	doUnary(c)
}
func doUnary(c addpb.AddServiceClient) {
	fmt.Println("starting to do add...")
	req := &addpb.AddRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Add:%v", err)
	}
	log.Printf("Response from Add:%v", res.Res)
}
