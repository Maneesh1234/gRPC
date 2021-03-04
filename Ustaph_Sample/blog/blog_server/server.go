// FOR GENERATING PROTO PB FILE//
//protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"gopkg.in/mgo.v2/bson"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"gRpc/Udemy/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Create blog request")
	//PARSE THE DATA
	blog := req.GetBlog()

	//MAP IT BLOG ITEM
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	//INSERT DATA THROUGH MONGODB DRIVER
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil

}

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	fmt.Println("Read blog request")

	blogID := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &blogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: dataToBlogPb(data),
	}, nil
}

func dataToBlogPb(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}
}



func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Blog Service Started")

	// DATABASE NAME=mydb , TABLE NAME=blog
	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	// client.Disconnect(context.TODO())
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}
	// Second step : closing the listener
	fmt.Println("Closing the listener")
	if err := lis.Close(); err != nil {
		log.Fatalf("Error on closing the listener : %v", err)
	}
	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
