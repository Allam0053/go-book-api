package book

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Calmantara/go-dts-user/module/model"
	"github.com/Calmantara/go-dts-user/module/service/book"
	"github.com/Calmantara/go-dts-user/pkg/response"
	"github.com/gin-gonic/gin"
)

type BookHdlImpl struct {
	bookSvc book.BookService
}

func NewBookHandler(bookSvc book.BookService) BookHandler {
	return &BookHdlImpl{
		bookSvc: bookSvc,
	}
}

func (u *BookHdlImpl) FindBookByIdHdl(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to find book",
			ErrorMsg: response.InvalidQuery,
		})
		return
	}
	// transform id string to uint64
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to find book",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	// call service
	book, err := u.bookSvc.FindBookByIdSvc(ctx, idUint)
	if err != nil {
		if err.Error() == "book is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to find book",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to find book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success find book",
		Data:    book,
	})
}

func (u *BookHdlImpl) FindAllBooksHdl(ctx *gin.Context) {
	books, err := u.bookSvc.FindAllBooksSvc(ctx)
	if err != nil {
		// bad code, should be wrapped in other package
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to get books",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success get books",
		Data:    books,
	})
}

func (u *BookHdlImpl) InsertBookHdl(ctx *gin.Context) {
	// mendapatkan body
	var usrIn model.Book

	if err := ctx.Bind(&usrIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to insert book",
			ErrorMsg: response.InvalidBody,
		})
		return
	}

	// validate name and email
	// if usrIn.Email == "" || usrIn.Name == "" {
	// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
	// 		Message:  "failed to insert book",
	// 		ErrorMsg: response.InvalidParam,
	// 	})
	// 	return
	// }

	insertedBook, err := u.bookSvc.InsertBookSvc(ctx, usrIn)
	if err != nil {
		// bad code, should be wrapped in other package
		if err.Error() == "error duplication email" {
			ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Message:  "failed to insert book",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to insert book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success create book",
		Data:    insertedBook,
	})
}

func (u *BookHdlImpl) UpdateBookHdl(ctx *gin.Context) {
	idUint, err := u.getIdFromParam(ctx)
	if err != nil {
		return
	}
	// binding payload
	var usrIn model.Book
	if err := ctx.Bind(&usrIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update book",
			ErrorMsg: response.InvalidBody,
		})
		return
	}
	usrIn.Id = idUint

	// validate name
	// if usrIn.Name == "" {
	// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
	// 		Message:  "failed to update book",
	// 		ErrorMsg: response.InvalidParam,
	// 	})
	// 	return
	// }

	if err := u.bookSvc.UpdateBookSvc(ctx, usrIn); err != nil {
		if err.Error() == "book is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to update book",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to update book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success update book",
	})
}

func (u *BookHdlImpl) DeleteBookByIdHdl(ctx *gin.Context) {
	idUint, err := u.getIdFromParam(ctx)
	if err != nil {
		return
	}
	deletedBook, err := u.bookSvc.DeleteBookByIdSvc(ctx, idUint)
	if err != nil {
		if err.Error() == "book is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to delete book",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to delete book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success delete book",
		Data:    deletedBook,
	})
}

func (u *BookHdlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")
	if id == "" {
		err = errors.New("failed id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update book",
			ErrorMsg: response.InvalidParam,
		})
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update book",
			ErrorMsg: response.InvalidParam,
		})
		return
	}
	return
}
