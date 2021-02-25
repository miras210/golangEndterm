package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	//	"time"

	"com.grpc.tleu/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewCalculatorServiceClient(conn)
	//doManyTimesFromServer(c)
	doAverage(c)
}



func doManyTimesFromServer(c greetpb.CalculatorServiceClient) {
	ctx := context.Background()
	req := &greetpb.PrimeRequest{Number: 120}

	stream, err := c.PrimeNumberDecomposition(ctx, req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from GreetManyTimes RPC %v", err)
		}
		log.Printf("response from PrimeNumberDecomposition: %v \n", res.GetNumber())
	}

}

func doAverage(c greetpb.CalculatorServiceClient) {

	requests := []*greetpb.AverageRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	ctx := context.Background()
	stream, err := c.ComputeAverage(ctx)
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("ComputeAverage Response: %v\n", res.GetNumber())
}