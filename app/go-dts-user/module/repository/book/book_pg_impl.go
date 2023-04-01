package book

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Calmantara/go-dts-user/module/model"
)

type BookPgRepoImpl struct {
	db *sql.DB
}

func NewBookPgRepo(db *sql.DB) BookRepo {
	return &BookPgRepoImpl{
		db: db,
	}
}

func (u *BookPgRepoImpl) FindBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	query := `
		SELECT 
			id, 
			name_book, 
			author
		FROM books u
		WHERE u.id = $1
			AND deleted_at is null;
	`
	// prepare untuk checking error di query
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx, bookId)
	if err != nil {
		if strings.Contains(err.Error(), "books_email_key") {
			err = errors.New("error duplication email")
		}
		return
	}
	defer rows.Close()

	for rows.Next() {
		// SCAN -> binding data ke golang struct
		if err = rows.Scan(&book.Id, &book.NameBook, &book.Author); err != nil {
			return
		}
	}

	if book.Id <= 0 {
		err = errors.New("book is not found")
	}

	return
}

func (u *BookPgRepoImpl) FindAllBooks(ctx context.Context) (books []model.Book, err error) {
	query := `
		SELECT 
			id, 
			name_book, 
			author
		FROM books u
		WHERE deleted_at is null
		ORDER BY created_at ASC;
	`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		if err = rows.Scan(&book.Id, &book.NameBook, &book.Author); err != nil {
			return
		}
		books = append(books, book)
	}
	return
}

func (u *BookPgRepoImpl) InsertBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	query := `
		INSERT INTO books
				(name_book, author)
		VALUES
			($1, $2 )
		RETURNING 
			id, name_book, author;
	`
	// RETURNING akan mengeluarkan
	// affected rows dari hasil query kita
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx,
		bookIn.NameBook,
		bookIn.Author)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&book.Id, &book.NameBook, &book.Author); err != nil {
			return
		}
	}
	return
}

func (u *BookPgRepoImpl) UpdateBook(ctx context.Context, bookIn model.Book) (err error) {
	query := `
		UPDATE books
		SET
			name_book  = $2,
			author = $3
		WHERE id = $1
			AND deleted_at is null;
	`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		bookIn.Id,
		bookIn.NameBook,
		bookIn.Author)
	if err != nil {
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affected <= 0 {
		err = errors.New("book is not found")
		return
	}
	return
}

func (u *BookPgRepoImpl) DeleteBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	query := `
		UPDATE books
		SET
			deleted_at = now()
		WHERE id = $1 
			AND deleted_at is null
		RETURNING 
			id, name_book, author;
	`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx,
		bookId)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&book.Id, &book.NameBook, &book.Author); err != nil {
			return
		}
	}
	if book.Id <= 0 {
		err = errors.New("book is not found")
	}
	return
}
