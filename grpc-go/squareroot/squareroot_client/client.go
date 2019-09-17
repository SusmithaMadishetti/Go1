package main

import (
	"context"
	"fmt"
	"log"

	squareroot "../squarerootpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello,I'm client!!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect:%v", err)
	}
	defer cc.Close()
	c := squareroot.NewSquarerootServiceClient(cc)
	doErrorUnary(c)
}
func doErrorUnary(c squareroot.SquarerootServiceClient) {
	fmt.Println("starting to do squareroot...")
	//number := int32(10)
	//correct call

	doErrorCall(c, 10)

	//error call
	doErrorCall(c, -2)

}
func doErrorCall(c squareroot.SquarerootServiceClient, n int32) {
	res, err := c.Squareroot(context.Background(), &squareroot.SquarerootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			//actual error from gRPC
			fmt.Printf("error msg from server:%v\n", respErr.Message())
			fmt.Printf("error code from server:%v\n", respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("we probably sent a negative number!\n")
				return
			}
		} else {
			log.Fatalf("big error calling Squareroot:%v", err)
			return
		}
	}
	fmt.Printf("Result of square root of %v : %v \n", n, res.GetResult())
}
