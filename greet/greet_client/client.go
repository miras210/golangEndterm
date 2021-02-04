package main

import (
	"context"
	"fmt"
	"gitlab.com/tleuzhan13/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Println("Hello I'm a client")

	conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v",err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient)  {
	ctx := context.Background()
	request := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		FirstName: "Tleuzhan",
		LastName:  "Mukatayev",
	}}

	responce,err := c.Greet(ctx,request)
	if err != nil {
		log.Fatalf("error while calling Greet RPC $v",err)
	}
	log.Printf("response from Greet:%v",responce.Result)
}