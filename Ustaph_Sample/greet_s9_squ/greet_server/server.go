package main

////COMMAND FOR CODE GENERATION
//protoc greet_s9_squ/greetpb/greet.proto --go_out=plugins=grpc:.

//FOR RUNNING
//cd "greet_s9_squ"
//go run greet_server\server.go

// ANOTHER TERMINAL
//go run greet_client\client.go
import (
	"context"
	"fmt"
	"gRpc/Udemy/greet_s9_squ/greetpb"
	"log"
	"math"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

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

// SEE LINE 849 IN GREET.PB.GO
//FindMaximum(CalculatorService_FindMaximumServer)
func (*server) FindMaximum(stream greetpb.CalculatorService_FindMaximumServer) error {
	// IN THIS STREAM HAVE stream.Recv() AND stream.Send( METHOD
	fmt.Printf("FindMaximum function was invoked with a streaming request\n")
	var m int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		Num := req.GetNum()
		//result := "Hello " + firstName + "! "

		if Num > m {
			m = Num
			sendErr := stream.Send(&greetpb.FindMaximumResponse{
				Max: m,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}

	}

}

//
func (*server) SquareRoot(ctx context.Context, req *greetpb.SquareRootRequest) (*greetpb.SquareRootResponse, error) {
	fmt.Println("Received SquareRoot RPC")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &greetpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
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
