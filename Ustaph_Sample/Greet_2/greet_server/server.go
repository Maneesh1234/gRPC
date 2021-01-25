package main

import (
	//"context"
	"fmt"
	//"io"
	"log"
	"net"
	//"strconv"
	//"time"

	//"google.golang.org/grpc/codes"
	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/status"

	"gRpc/Udemy/greet_2/greetpb"
	//"github.com/simplesteph/grpc-go-course/greet/greetpb"
	//C:\Go\src\gRpc\Udemy\greet_2\greetpb
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello world")

	//tcp connection
	// 50051 is the default port number for grpc
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// opts := []grpc.ServerOption{}
	// tls := false
	// if tls {
	// 	certFile := "ssl/server.crt"
	// 	keyFile := "ssl/server.pem"
	// 	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	// 	if sslErr != nil {
	// 		log.Fatalf("Failed loading certificates: %v", sslErr)
	// 		return
	// 	}
	// 	opts = append(opts, grpc.Creds(creds))
	// }

	//CREATE A GRPC SERVER
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	//BIND THE PORT NUMBER  TO THE GRPC SERVER
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
