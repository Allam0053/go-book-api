package book

import (
	"github.com/Calmantara/go-dts-user/module/handler/book"
	"github.com/gin-gonic/gin"
)

func NewBookRouter(v1 *gin.RouterGroup, bookHdl book.BookHandler) {
	g := v1.Group("/book")

	// register all router
	g.GET("/all", bookHdl.FindAllBooksHdl)
	g.GET("", bookHdl.FindBookByIdHdl)
	g.POST("", bookHdl.InsertBookHdl)
	g.PUT("/:id", bookHdl.UpdateBookHdl)
	g.DELETE("/:id", bookHdl.DeleteBookByIdHdl)
}
