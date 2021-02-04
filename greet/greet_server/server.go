package main

import (
	"context"
	"fmt"
	"gitlab.com/tleuzhan13/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{
	greetpb.UnimplementedGreetServiceServer
}

func (s *Server) Greet(ctx context.Context,req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet function was invoked with %v \n",req)
	firstName := req.GetGreeting().GetFirstName()
	
	result := "Hello "+firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	
	return res,nil
}

func main()  {
	l,err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v",err)
	}

	s:=grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s,&Server{})

	if err := s.Serve(l);err!=nil {
		log.Fatalf("failed to serve:%v",err)
	}
}
