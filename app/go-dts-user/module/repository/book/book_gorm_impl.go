package book

import (
	"context"
	"errors"
	"log"

	"github.com/Calmantara/go-dts-user/config"
	"github.com/Calmantara/go-dts-user/module/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookGormRepoImpl struct {
	db *gorm.DB
}

func NewBookGormRepo(db *gorm.DB) BookRepo {
	bookRepo := &BookGormRepoImpl{
		db: db,
	}

	if config.Load.DataSource.Migrate {
		bookRepo.doMigration()
	}

	return &BookGormRepoImpl{
		db: db,
	}
}

func (u *BookGormRepoImpl) doMigration() (err error) {
	// create book table
	if err = u.db.AutoMigrate(&model.Book{}); err != nil {
		panic(err)
	}
	log.Println("successfully create book table")
	// create book photo table
	// if err = u.db.AutoMigrate(&model.BookPhoto{}); err != nil {
	// 	panic(err)
	// }
	log.Println("successfully create book photo table")
	return
}

func (u *BookGormRepoImpl) FindBookByIdEager(ctx context.Context, bookId uint64) (book model.Book, err error) {
	// eager from book to photo
	tx := u.db.
		Model(&model.Book{}).
		Preload("Photos").
		Where("id = ?", bookId).
		Find(&book)
	if err = tx.Error; err != nil {
		return
	}

	// eager from photo to book
	// var bookPhoto model.BookPhoto
	// tx = u.db.
	// 	Model(&model.BookPhoto{}).
	// 	Preload("BookDetail").
	// 	Where("id = ?", 1).
	// 	Find(&bookPhoto)
	// if err = tx.Error; err != nil {
	// 	return
	// }

	if book.Id <= 0 {
		err = errors.New("book is not found")
	}
	return
}

func (u *BookGormRepoImpl) FindBookByIdJoin(ctx context.Context, bookId uint64) (err error) {

	type BookCustom struct {
		Id       uint64 `json:"id" gorm:"column:id"`
		Name     string `json:"name" gorm:"column:name"`
		Email    string `json:"email" gorm:"column:email"`
		PhotoId  uint64 `json:"photo_id" gorm:"column:photo_id"`
		PhotoUrl string `json:"photo_url" gorm:"column:photo_url"`
	}

	// join with new struct type
	var books []BookCustom
	tx := u.db.
		Table("books").
		Select(`
			books.id as id, 
			book_photos.id as photo_id,
			books.name as name,
			books.email as email,
			book_photos.url as photo_url
		`).
		Joins(`JOIN book_photos 
			   ON books.id = book_photos.book_id
			   AND book_photos.deleted_at IS NULL`).
		Where("books.id = ?", bookId).
		Find(&books)
	if err = tx.Error; err != nil {
		return
	}

	// eager join
	var bookModel []model.Book
	tx = u.db.
		Joins(`Photo`).
		Where("books.id = ?", bookId).
		Find(&bookModel)
	if err = tx.Error; err != nil {
		return
	}

	return
}

func (u *BookGormRepoImpl) FindBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	tx := u.db.
		Model(&model.Book{}).
		Where("id = ?", bookId).
		Find(&book)

	if err = tx.Error; err != nil {
		return
	}

	if book.Id <= 0 {
		err = errors.New("book is not found")
	}

	return
}

func (u *BookGormRepoImpl) FindAllBooks(ctx context.Context) (books []model.Book, err error) {
	tx := u.db.
		Model(&model.Book{}).
		Find(&books).
		Order("created_at ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (u *BookGormRepoImpl) InsertBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	tx := u.db.
		Model(&model.Book{}).
		Create(&bookIn)

	if err = tx.Error; err != nil {
		return
	}

	return bookIn, err
}

func (u *BookGormRepoImpl) BulkInsertBook(ctx context.Context, bookIn []model.Book) (err error) {
	tx := u.db.
		Model(&model.Book{}).
		Create(&bookIn)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (u *BookGormRepoImpl) UpdateBook(ctx context.Context, bookIn model.Book) (err error) {
	tx := u.db.
		Model(&model.Book{}).
		Where("id = ?", bookIn.Id).
		Updates(&bookIn)

	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("book is not found")
		return
	}

	return
}

func (u *BookGormRepoImpl) DeleteBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	tx := u.db.
		Model(&model.Book{}).
		// clause to return data after delete
		Clauses(clause.Returning{}).
		Where("id = ?", bookId).
		Delete(&book)
	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("book is not found")
		return
	}
	return
}

func (u *BookGormRepoImpl) HardDeleteBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	tx := u.db.
		Unscoped().
		Model(&model.Book{}).
		Where("id = ?", bookId).
		Delete(&model.Book{})
	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("book is not found")
		return
	}
	return
}
