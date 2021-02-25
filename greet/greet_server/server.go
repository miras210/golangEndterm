package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"com.grpc.tleu/greet/greetpb"
	"google.golang.org/grpc"
)

//Server with embedded UnimplementedGreetServiceServer
type Server struct {
	greetpb.UnimplementedCalculatorServiceServer
}

func (s *Server) PrimeNumberDecomposition(req *greetpb.PrimeRequest, stream greetpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked with %v \n", req)
	num := req.GetNumber()
	var count int32 = 2
	for num != 1 {
		if num%count == 0 {
			res := &greetpb.PrimeResponse{Number: count}
			if err := stream.Send(res); err != nil {
				log.Fatalf("error while sending greet many times responses: %v", err.Error())
			}
			num = num / count
		} else {
			count += 1
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *Server) ComputeAverage(stream greetpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with \n")
	result := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&greetpb.AverageResponse{
				Number: result / count,
			})

		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		num := req.GetNumber()
		count += 1
		result += float64(num)
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
