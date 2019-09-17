package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"../greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("im client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doBiDiStreaming(c)
	doUnaryWithDeadline(c, 5*time.Second)
	doUnaryWithDeadline(c, 1*time.Second)

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
	fmt.Println("starting Bidirectional")
	//we create stream by invoking client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream:%v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "susmitha",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "karthik",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "madishetti",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "pondicherry",
			},
		},
	}

	waitc := make(chan struct{})
	//we send bunch of messages to the client(go routine)
	go func() {
		//function to send bunch of messages
		for _, req := range requests {
			fmt.Printf("sending message:%v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	//we receive bunch of messages from the client(go routine)
	go func() {
		//function to receive bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while reading:%v", err)
				break
			}
			fmt.Printf("received:%v\n", res.GetResult())
		}
		close(waitc)
	}()
	//block until evrything is done
	<-waitc

}
func doUnaryWithDeadline(c greetpb.GreetServiceClient, seconds time.Duration) {
	fmt.Println("starting to do UnaryWithDeadline RPC..")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Susmitha",
			LastName:  "Madishetti",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("timeout was hit!Deadline was exceeded")
			} else {
				fmt.Printf("enexpected error:%v", statusErr)
			}
		} else {
			log.Fatalf("error while callimg greetwith deadline Rpc:%v", err)
		}
	}
	log.Printf("Response from Greet:%v", res.Result)
}
