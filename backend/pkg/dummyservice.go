package pkg

import (
	"context"
	"fmt"
	sdk "git.gendocu.com/gendocu/LibraryApp.git/sdk/go"
	"github.com/golang/protobuf/ptypes/empty"
)

type DummyService struct {
	books []*sdk.Book
}

func NewDummyService() *DummyService {
	return &DummyService{
		books: []*sdk.Book{{
			Isbn:   "0-670-81302-8",
			Title:  "It",
			Author: &sdk.Author{
				FirstName: "Stephen",
				LastName: "King",
			},
		}},
	}
}
func (s *DummyService) ListBooks(ctx context.Context, empty *empty.Empty) (*sdk.ListBookResponse, error) {
	return &sdk.ListBookResponse{
		Books: s.books,
	}, nil
}

func (s *DummyService) DeleteBook(ctx context.Context, request *sdk.DeleteBookRequest) (*sdk.Book, error) {
	for ind, book := range s.books {
		if book.Isbn == request.GetIsbn() || book.Title == request.GetTitle(){
			//removing the book
			s.books[ind] = s.books[len(s.books)-1]
			s.books = s.books[:len(s.books)-1]
			return book, nil
		}
	}
	return nil, fmt.Errorf("book not found %+v", request)
}

func (s *DummyService) CreateBook(ctx context.Context, book *sdk.Book) (*sdk.Book, error) {
	s.books = append(s.books, book)
	return book, nil
}

