package book

import (
	"context"

	"github.com/Calmantara/go-dts-user/module/model"
)

type BookService interface {
	FindBookByIdSvc(ctx context.Context, bookId uint64) (book model.Book, err error)
	FindAllBooksSvc(ctx context.Context) (books []model.Book, err error)
	InsertBookSvc(ctx context.Context, bookIn model.Book) (book model.Book, err error)
	UpdateBookSvc(ctx context.Context, bookIn model.Book) (err error)
	DeleteBookByIdSvc(ctx context.Context, bookId uint64) (book model.Book, err error)
}
