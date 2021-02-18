package main

////COMMAND FOR CODE GENERATION
//protoc greet_s7_avg/greetpb/greet.proto --go_out=plugins=grpc:.

//FOR RUNNING
//cd "greet_s7_avg"
//go run greet_server\server.go

// ANOTHER TERMINAL
//go run greet_client\client.go
import (
	//"context"
	"fmt"
	"gRpc/Udemy/greet_s7_avg/greetpb"
	"log"
	"net"
	//"strconv"
	"io"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

// // SEE LINE 330 IN Greet.PB.GO
// func (*server) Greet(ctx context.Context, req *Greetpb.GreetRequest) (*Greetpb.GreetResponse, error) {
// 	fmt.Printf("Greet function was invoked with %v\n", req)
// 	// SEE LINE 68 IN Greet.PB.GO
// 	firstName := req.GetGreeting().GetFirstName()
// 	result := "Hello " + firstName
// 	res := &Greetpb.GreetResponse{
// 		Result: result,
// 	}
// 	return res, nil
// }

// // SEE LINE 505 IN Greet.PB.GO
// func (*server) GreetManyTimes(req *Greetpb.GreetManyTimesRequest, stream Greetpb.GreetService_GreetManyTimesServer) error {
// 	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
// 	firstName := req.GetGreeting().GetFirstName()
// 	for i := 0; i < 10; i++ {
// 		result := "Hello " + firstName + " number " + strconv.Itoa(i)
// 		res := &Greetpb.GreetManytimesResponse{
// 			Result: result,
// 		}
// 		stream.Send(res)
// 		time.Sleep(1000 * time.Millisecond)
// 	}
// 	return nil
// }

// SEE LINE 350 IN Greet.PB.GO
func (*server) PrimeManyTimes(req *greetpb.PrimeNumberRequest, stream greetpb.CalculatorService_PrimeManyTimesServer) error {
	fmt.Printf("PrimeManyTimes function was invoked with %v\n", req)
	num := req.GetPrimeNumberDecomposition().GetNum()
	var k int64 = 2
	for num > 1 {

		//N = 210
		//while num > 1:
		if num%k == 0 { // if k evenly divides into N
			//print k      // this is a factor
			result := k
			res := &greetpb.PrimeNumberResponse{
				Result: result,
			}
			stream.Send(res)
			num = num / k
		} else { // divide N by k so that we have the rest of the number left.
			k = k + 1
		}
		// result := "Hello " + firstName + " number " + strconv.Itoa(i)
		// res := &Greetpb.GreetManytimesResponse{
		// 	Result: result,
		// }
		// stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

// SEE LINE 531 IN GREET.PB.GO

func (*server) ComputeAverage(stream greetpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a client streaming request\n")
	avg := int64(0)
	i := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// SEE LINE 743 IN GREET.PB.GO
			// WE HAVE TO FINISH THE CLIENT STREAM
			avg = avg / i
			return stream.SendAndClose(&greetpb.ComputeAverageResponse{
				Avg: avg,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		// SEE LINE 208 IN GREET.PB.GO
		num := req.GetNum()
		avg += num
		i += 1
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
	greetpb.RegisterCalculatorServiceServer(s, &server{})

	//BIND THE PORT NUMBER  TO THE GRPC SERVER
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
