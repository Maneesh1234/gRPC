package main

// FOR RUNNING SERVER
// CD UDEMY
//cd "greet_s7_cli"
//go run greet_server\server.go

////COMMAND FOR CODE GENERATION
//protoc greet_s7_cli/greetpb/greet.proto --go_out=plugins=grpc:.

import (
	"context"
	"fmt"
	"gRpc/Udemy/greet_s7_cli/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

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

// SEE LINE 505 IN GREET.PB.GO
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManytimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

// SEE LINE 691 IN GREET.PB.GO
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a client streaming request\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// SEE LINE 743 IN GREET.PB.GO
			// WE HAVE TO FINISH THE CLIENT STREAM
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
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
