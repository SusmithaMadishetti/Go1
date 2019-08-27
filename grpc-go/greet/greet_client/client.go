package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"../greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("im client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
	doBiDiStreaming(c)
}
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do Unary RPC..")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Susmitha",
			LastName:  "Madishetti",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while callimg greet Rpc:%v", err)

	}
	log.Printf("Response from Greet:%v", res.Result)
}
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do ServerStreaming...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Susmitha",
			LastName:  "Madishetti",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes:%v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading Stream:%v", err)

		}
		log.Printf("Response from Greet:%v", msg.GetResult())
	}
}
func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("streaming client Streaming rpc..")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "susmitha",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "karthik",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "madishetti",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "pondicherry",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling longGreet:%v", err)
	}
	//iterate our slice and each message individually
	for _, req := range requests {
		fmt.Printf("sending req :%v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving:%v", err)
	}
	fmt.Printf("LongGreet Response:%v\n", res)
}
func doBiDiStreaming(c greetpb.GreetServiceClient) {

}
