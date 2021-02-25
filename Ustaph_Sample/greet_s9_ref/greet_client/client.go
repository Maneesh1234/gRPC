package main

import (
	"context"
	"fmt"
	"gRpc/Udemy/greet_s9_ref/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

//FOR REFLECTION COMMAND
//evans
//STRAT EVANS IN REFLECTION MODE
//evans -p 50051 -r

//show package
//show service
//show message
//service CalculaatorService
//call ComputeAverage

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

//// SEE LINE 435 IN GREET.PB.GO
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

//// SEE LINE 435 IN GREET.PB.GO
func doClientStreaming(c greetpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streaming for average calculation RPC...")

	// (1 ) WAY
	// requests := []*greetpb.ComputeAverageRequest{
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 1,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 2,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 3,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 4,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 5,
	// 	},
	// }

	// // SEE LINE 482 IN GREET.PB.GO
	// stream, err := c.ComputeAverage(context.Background())
	// if err != nil {
	// 	log.Fatalf("error while calling ComputeAverage: %v", err)
	// }

	// // we iterate over our slice and send each message individually
	// for _, req := range requests {
	// 	fmt.Printf("Sending req: %v\n", req)
	// 	// SEE LINE 501 IN GREET.PB.GO
	// 	stream.Send(req)
	// 	time.Sleep(1000 * time.Millisecond)
	// }

	//(2 ) WAY
	numbers := []int64{1, 20, 3, 4, 5}
	// requests := []*greetpb.ComputeAverageRequest{
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 1,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 2,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 3,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 4,
	// 	},
	// 	&greetpb.ComputeAverageRequest{

	// 		Num: 5,
	// 	},
	// }

	// SEE LINE 482 IN GREET.PB.GO
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling ComputeAverage: %v", err)
	}

	// we iterate over our slice and send each message individually
	for _, num := range numbers {
		fmt.Printf("Sending num: %v\n", num)
		// SEE LINE 501 IN GREET.PB.GO
		stream.Send(&greetpb.ComputeAverageRequest{

			Num: num,
		})
		time.Sleep(1000 * time.Millisecond)
	}

	// SEE LINE 505  ,493 IN GREET.PB.GO
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ComputeAverage: %v", err)
	}
	//LINE NUMBER 255
	fmt.Printf("ComputeAverage Response: %v\n", res.GetAvg())

}

func doBiDiStreaming(c greetpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")
	numbers := []int64{1, 20, 3, 4, 25, 7, 65, 80, 90}
	// we create a stream by invoking the client
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	// requests := []*greetpb.FindMaximumRequest{
	// 	&greetpb.GreetEveryoneRequest{
	// 		Greeting: &greetpb.Greeting{
	// 			FirstName: "Maneesh",
	// 		},
	// 	},
	// 	&greetpb.GreetEveryoneRequest{
	// 		Greeting: &greetpb.Greeting{
	// 			FirstName: "tom",
	// 		},
	// 	},
	// 	&greetpb.GreetEveryoneRequest{
	// 		Greeting: &greetpb.Greeting{
	// 			FirstName: "jerry",
	// 		},
	// 	},
	// 	&greetpb.GreetEveryoneRequest{
	// 		Greeting: &greetpb.Greeting{
	// 			FirstName: "raju",
	// 		},
	// 	},
	// 	&greetpb.GreetEveryoneRequest{
	// 		Greeting: &greetpb.Greeting{
	// 			FirstName: "rakesh",
	// 		},
	// 	},
	// }

	// // we iterate over our slice and send each message individually
	// for _, num := range numbers {
	// 	fmt.Printf("Sending num: %v\n", num)
	// 	// SEE LINE 501 IN GREET.PB.GO
	// 	stream.Send(&greetpb.ComputeAverageRequest{

	// 		Num: num,
	// 	})
	// 	time.Sleep(1000 * time.Millisecond)
	// }

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, num := range numbers {
			fmt.Printf("Sending num: %v\n", num)
			// SEE LINE 501 IN GREET.PB.GO
			stream.Send(&greetpb.FindMaximumRequest{

				Num: num,
			})
			time.Sleep(2000 * time.Millisecond)
		}
		// THIS DENOTES CLIENT END SENDING DATA
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			//SEE LINE 349
			fmt.Printf("Received: %v\n", res.GetMax())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}

func doErrorUnary(c greetpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 16)

	// error call
	doErrorCall(c, -5)
}

func doErrorCall(c greetpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &greetpb.SquareRootRequest{Number: n})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", n, res.GetNumberRoot())
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
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doBiDiStreaming(c)
	doErrorUnary(c)

}
