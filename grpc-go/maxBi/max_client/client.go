package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"../maxpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnt connect:%v", err)
	}
	defer conn.Close()
	c := maxpb.NewMaxServiceClient(conn)
	doBidi(c)
}
func doBidi(c maxpb.MaxServiceClient) {
	fmt.Println("starting  bidirectional streaming..")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream and calling Max function :%v", err)
	}
	waitc := make(chan struct{})

	//send go routine
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			fmt.Printf("Sending number:%v\n", number)
			stream.Send(&maxpb.MaxRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//receive goroutine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("problem while reading server stream:%v", err)
				break
			}
			resu := res.GetResult()
			fmt.Printf("received a new maximum of  %v \n", resu)
		}
		close(waitc)
	}()
	<-waitc

}
