package book

import (
	"context"
	"errors"

	"github.com/Calmantara/go-dts-user/config"
	"github.com/Calmantara/go-dts-user/module/model"
)

type BookMapImpl struct {
	dataStore  config.DataStore
	cacheEmail map[string]bool
}

func NewBookMap(dataStore config.DataStore) BookRepo {
	cache := make(map[string]bool)
	return &BookMapImpl{
		dataStore:  dataStore,
		cacheEmail: cache,
	}
}

func (um *BookMapImpl) FindBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	book, ok := um.dataStore.BookData[bookId]
	if !ok || book.Delete {
		err = errors.New("book is not found")
	}
	return
}

func (um *BookMapImpl) FindAllBooks(ctx context.Context) (books []model.Book, err error) {
	for _, usr := range um.dataStore.BookData {
		if !usr.Delete {
			books = append(books, usr)
		}
	}
	return
}

func (um *BookMapImpl) InsertBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	// if um.cacheEmail[bookIn.Email] {
	// 	err = errors.New("error duplication email")
	// 	return
	// }
	// insert
	bookIn.Id = uint64(len(um.dataStore.BookData) + 1)
	um.dataStore.BookData[bookIn.Id] = bookIn // store book
	// um.cacheEmail[bookIn.Email] = true        // store email
	return bookIn, err
}

func (um *BookMapImpl) UpdateBook(ctx context.Context, bookIn model.Book) (err error) {
	// update
	book, err := um.FindBookById(ctx, bookIn.Id)
	if err != nil && book.Id > 0 {
		return err
	}
	// delete(um.cacheEmail, book.Email)         // delete previous email cache
	um.dataStore.BookData[bookIn.Id] = bookIn // update book map
	// um.cacheEmail[bookIn.Email] = true        // store email cache
	return err
}

func (um *BookMapImpl) DeleteBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {

	return
}
