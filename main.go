package main

import (
	"log"
	"net"

	"github.com/Jayprakash1234/bookservice/database"
	"github.com/Jayprakash1234/bookservice/pb"
	"github.com/Jayprakash1234/bookservice/server"

	"google.golang.org/grpc"
)

func main() {
	/* router := gin.Default()

	database.InitDB()

	router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.AddBook)

	router.Run(":8080") */

	database.InitDB()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, &server.BookServer{})

	log.Println("🚀 gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
