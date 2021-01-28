package main

import (
	"context"
	"fmt"

	"log"
	"net"

	"gRpc/Udemy/greet_4_20/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

// SEE LINE 330 IN GREET.PB.GO
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	// SEE LINE 68 IN GREET.PB.GO
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	//tcp connection
	// 50051 is the default port number for grpc
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//CREATE A GRPC SERVER
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	//BIND THE PORT NUMBER  TO THE GRPC SERVER
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
