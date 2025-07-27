package server

import (
	"context"

	"github.com/Jayprakash1234/bookservice/database"
	"github.com/Jayprakash1234/bookservice/models"
	"github.com/Jayprakash1234/bookservice/pb"
)

type BookServer struct {
	pb.UnimplementedBookServiceServer
}

func (s *BookServer) GetAllBooks(ctx context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	var books []models.Book
	database.DB.Find(&books)

	var pbBooks []*pb.Book
	for _, b := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:     b.ID,
			Title:  b.Title,
			Author: b.Author,
		})
	}

	return &pb.GetAllBooksResponse{Books: pbBooks}, nil
}

func (s *BookServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	book := models.Book{
		Title:  req.GetTitle(),
		Author: req.GetAuthor(),
	}

	database.DB.Create(&book)

	return &pb.AddBookResponse{
		Book: &pb.Book{
			Id:     book.ID,
			Title:  book.Title,
			Author: book.Author,
		},
	}, nil
}
