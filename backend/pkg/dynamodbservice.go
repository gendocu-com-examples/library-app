package pkg

import (
	"context"
	sdk "git.gendocu.com/gendocu/LibraryApp.git/sdk/go"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/guregu/dynamo"
)

type DynamoDBService struct {
	t dynamo.Table
}

func NewDynamoDBService() *DynamoDBService {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	db := dynamo.New(sess)
	return &DynamoDBService{
		t: db.Table("ExampleLibraryApp"),
	}
}
func (s *DynamoDBService) ListBooks(ctx context.Context, empty *empty.Empty) (*sdk.ListBookResponse, error) {
	var res []*sdk.Book
	if err := s.t.Scan().All(&res); err != nil {
		return nil, err
	}
	return &sdk.ListBookResponse{
		Books: res,
	}, nil
}

func (s *DynamoDBService) DeleteBook(ctx context.Context, request *sdk.DeleteBookRequest) (*sdk.Book, error) {
	var res *sdk.Book
	if err := s.t.Get("Isbn", request.GetIsbn()).OneWithContext(ctx, &res); err != nil {
		return nil, err
	}
	err := s.t.Delete("Isbn", request.GetIsbn()).RunWithContext(ctx)
	return res, err
}

func (s *DynamoDBService) CreateBook(ctx context.Context, book *sdk.Book) (*sdk.Book, error) {
	err := s.t.Put(book).RunWithContext(ctx)
	return book, err
}

