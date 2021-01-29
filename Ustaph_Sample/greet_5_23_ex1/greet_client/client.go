package main

import (
	"context"
	"fmt"
	//"io"
	"log"
	//"time"

	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/credentials"
	"gRpc/Udemy/greet_5_23_ex1/greetpb"

	//"github.com/simplesteph/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/status"
)

// SEE LINE 305 IN GREET.PB.GO
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstNum:  10,
			SecondNum: 3,
		},
	}
	// SEE LINE 307 IN GREET.PB.GO
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func main() {

	fmt.Println("Hello I'm a client")


	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	doUnary(c)
	
}

