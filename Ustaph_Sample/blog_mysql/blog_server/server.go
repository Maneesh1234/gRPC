
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"gRpc/Comp_grpc/blog_mysql/blogpb"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
	"google.golang.org/grpc"
)


type server struct {
	db *sql.DB
}

type blogItem struct {

	AuthorID int64 `db:"author_id" json:"author_id"`
	Content  string `db:"content" json:"content"`
	Title     string    `db:"title" json:"title"`
	
}

// connect returns SQL database connection from the pool
func (s *server) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
func NewServer(db *sql.DB) blogpb.BlogServiceServer {
	return &server{db: db}
}

func (s *server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Create blog request")

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	//PARSE THE DATA
	blog := req.GetBlog()

	//MAP IT BLOG ITEM
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	//INSERT DATA THROUGH Mysql
	// insert blog entity data
	_, err = c.ExecContext(ctx, "INSERT INTO blog(`author_id`, `content`, `title`) VALUES(?, ?, ?)",
		data.AuthorID, data.Content,data.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Blog-> "+err.Error())
	}
	

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil

}



func dataToBlogPb(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}
}


func (s *server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	fmt.Println("Read blog request")

	
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	authorId:=req.GetAuthorId()
	data := &blogItem{}
	// query blog by ID
	rows, err := c.QueryContext(ctx, "SELECT `author_id`, `content`, `title` FROM blog WHERE `author_id`=?",
		authorId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from blog-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from blog-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("blog with author_id='%d' is not found",
			authorId))
	}

	
	if err := rows.Scan(&data.AuthorID, &data.Content, &data.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from blog row-> "+err.Error())
	}
	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple blog rows with 	authorId='%d'",
		authorId))
	}
	

	return &blogpb.ReadBlogResponse{
		Blog: dataToBlogPb(data),
	}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("Update blog request")
	
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	b := req.GetBlog()
	authorId:=b.GetAuthorId()
	data := &blogItem{}
	// update blog
	res, err := c.ExecContext(ctx, "UPDATE blog SET `content`=?, `title`=? WHERE `author_id`=?",
		b.GetContent(), b.GetTitle(), authorId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update blog-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("blog with authorId='%d' is not found",
		authorId))
	}
	// we update our internal struct
	data.AuthorID = b.GetAuthorId()
	data.Content = b.GetContent()
	data.Title = b.GetTitle()
	
	return &blogpb.UpdateBlogResponse{
		Blog: dataToBlogPb(data),
	}, nil

}
func (s *server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("Delete blog request")
	
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	authorId:=req.GetAuthorId()
	
	// delete blog
	res, err := c.ExecContext(ctx, "DELETE FROM blog WHERE `author_id`=?", authorId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete blog-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("blog with authorId='%d' is not found",
		authorId))
	}
	

	return &blogpb.DeleteBlogResponse{AuthorId: req.GetAuthorId()}, nil
}
func (s *server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {
	fmt.Println("List blog request")

	
	// get SQL connection from pool
	c, err := s.connect(context.Background())
	if err != nil {
		return err
	}
	defer c.Close()
	// get blog list
	rows, err := c.QueryContext(context.Background(), "SELECT  `author_id`, `content`, `title`  FROM blog")
	if err != nil {
		return status.Error(codes.Unknown, "failed to select from blog-> "+err.Error())
	}
	defer rows.Close()

	
	for rows.Next() {
		data := &blogItem{}
		if err := rows.Scan(&data.AuthorID, &data.Content, &data.Title); err != nil {
			return status.Error(codes.Unknown, "failed to retrieve field values from blog row-> "+err.Error())
		}
		
		
		stream.Send(&blogpb.ListBlogResponse{Blog: dataToBlogPb(data)})
	}
	if err := rows.Err(); err != nil {
		return  status.Error(codes.Unknown, "failed to retrieve data from blog-> "+err.Error())
	}
	

	return nil
}
func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	 db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/blog_pb")
	if err != nil {
		// fmt.Print(err.Error())
		fmt.Println("Error creating DB:", err)
		fmt.Println("To verify, db is:", db)
	}
	defer db.Close()
	fmt.Println("Successfully  Connected to MYSQl")
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	
	seri := NewServer(db)

	
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, seri)
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
