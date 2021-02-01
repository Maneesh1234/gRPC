package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gRpc/Udemy/greet_s6_pri/greetpb"

	"google.golang.org/grpc"
)

// // SEE LINE 305 IN GREET.PB.GO
// func doUnary(c greetpb.GreetServiceClient) {
// 	fmt.Println("Starting to do a Unary RPC...")
// 	req := &greetpb.GreetRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "MANEESH",
// 			LastName:  "KUMAR",
// 		},
// 	}
// 	// SEE LINE 307 IN GREET.PB.GO
// 	res, err := c.Greet(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("error while calling Greet RPC: %v", err)
// 	}
// 	log.Printf("Response from Greet: %v", res.Result)
// }

//// SEE LINE 307 IN GREET.PB.GO
func doServerStreaming(c greetpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streaming for prime decomposition RPC...")

	req := &greetpb.PrimeNumberRequest{
		PrimeNumberDecomposition: &greetpb.PrimeNumberDecomposition{
			Num: 120,
		},
	}

	resStream, err := c.PrimeManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeManyTimes RPC: %v", err)
	}
	for {
		// SEE LINE 484 IN GREET.PB.GO
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func main() {

	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	// SEE LINE 311 IN GREET.PB.GO
	c := greetpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)
	doServerStreaming(c)

}
