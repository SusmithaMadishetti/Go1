package main

import (
	"context"
	"fmt"
	"log"

	"../averagepb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnt conncet:%v", err)
	}
	defer conn.Close()
	c := averagepb.NewAverageServiceClient(conn)
	doaverageCS(c)
}
func doaverageCS(c averagepb.AverageServiceClient) {
	fmt.Println("strating average:\n")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while opeming stream:%v", err)
	}
	numbers := []int64{3, 5, 9, 54, 23}
	for _, number := range numbers {
		fmt.Printf("senidng number:%v\n", number)
		stream.Send(&averagepb.AverageRequest{
			Number: number,
		})

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("erroe while receiving erroe:%v", err)
	}
	fmt.Printf("Averageis :%v\n", res.GetResult())

}
