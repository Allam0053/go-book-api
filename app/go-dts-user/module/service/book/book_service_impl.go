package book

import (
	"context"
	"log"

	"github.com/Calmantara/go-dts-user/module/model"
	"github.com/Calmantara/go-dts-user/module/repository/book"
)

type BookSvcImpl struct {
	bookRepo book.BookRepo
}

func NewBookSvc(bookRepo book.BookRepo) BookService {
	return &BookSvcImpl{
		bookRepo: bookRepo,
	}
}

func (u *BookSvcImpl) FindBookByIdSvc(ctx context.Context, bookId uint64) (book model.Book, err error) {
	log.Printf("[INFO] %T FindBookById invoked\n", u)
	if book, err = u.bookRepo.FindBookById(ctx, bookId); err != nil {
		log.Printf("[ERROR] error FindBookById :%v\n", err)
	}
	return
}

func (u *BookSvcImpl) FindAllBooksSvc(ctx context.Context) (books []model.Book, err error) {
	log.Printf("[INFO] %T FindAllBooks invoked\n", u)
	if books, err = u.bookRepo.FindAllBooks(ctx); err != nil {
		log.Printf("[ERROR] error FindAllBooks :%v\n", err)
	}
	return
}

func (u *BookSvcImpl) InsertBookSvc(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	log.Printf("[INFO] %T InsertBook invoked\n", u)
	if book, err = u.bookRepo.InsertBook(ctx, bookIn); err != nil {
		log.Printf("[ERROR] error InsertBook :%v\n", err)
	}
	return
}

func (u *BookSvcImpl) UpdateBookSvc(ctx context.Context, bookIn model.Book) (err error) {
	log.Printf("[INFO] %T UpdateBook invoked\n", u)
	if err = u.bookRepo.UpdateBook(ctx, bookIn); err != nil {
		log.Printf("[ERROR] error InsertBook :%v\n", err)
	}
	return
}

func (u *BookSvcImpl) DeleteBookByIdSvc(ctx context.Context, bookId uint64) (deletedBook model.Book, err error) {
	log.Printf("[INFO] %T DeleteBookById invoked\n", u)
	if deletedBook, err = u.bookRepo.DeleteBookById(ctx, bookId); err != nil {
		log.Printf("[ERROR] error DeleteBookById :%v\n", err)
	}
	return
}
