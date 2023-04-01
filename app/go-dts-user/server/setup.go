package server

import (
	"context"

	"github.com/Calmantara/go-dts-user/config"
	bookhdl "github.com/Calmantara/go-dts-user/module/handler/book"
	userhdl "github.com/Calmantara/go-dts-user/module/handler/user"
	bookrepo "github.com/Calmantara/go-dts-user/module/repository/book"
	userrepo "github.com/Calmantara/go-dts-user/module/repository/user"
	booksvc "github.com/Calmantara/go-dts-user/module/service/book"
	usersvc "github.com/Calmantara/go-dts-user/module/service/user"
	c "github.com/Calmantara/go-dts-user/pkg/context"
	"github.com/Calmantara/go-dts-user/pkg/logger"
)

type handlers struct {
	userHdl userhdl.UserHandler
	bookHdl bookhdl.BookHandler
}

func initDI() handlers {
	ctx, _ := c.GetCorrelationID(context.Background())

	logger.Info(ctx, "setup repository")
	dataStore := config.ConnectDataStore()
	userRepo := userrepo.NewUserMap(dataStore)
	bookRepo := bookrepo.NewBookMap(dataStore)

	switch config.Load.DataSource.Mode {
	case config.MODE_GORM:
		pgConn := config.NewPostgresConn()
		userRepo = userrepo.NewUserPgRepo(pgConn)
		bookRepo = bookrepo.NewBookPgRepo(pgConn)
	case config.MODE_PG:
		pgConn := config.NewPostgresConn()
		userRepo = userrepo.NewUserPgRepo(pgConn)
		bookRepo = bookrepo.NewBookPgRepo(pgConn)
	default:
		pgConn := config.NewPostgresConn()
		userRepo = userrepo.NewUserPgRepo(pgConn)
		bookRepo = bookrepo.NewBookPgRepo(pgConn)
	}

	logger.Info(ctx, "setup service")
	userSvc := usersvc.NewUserSvc(userRepo)
	bookSvc := booksvc.NewBookSvc(bookRepo)

	logger.Info(ctx, "setup handler")
	userHdl := userhdl.NewUserHandler(userSvc)
	bookHdl := bookhdl.NewBookHandler(bookSvc)

	return handlers{
		userHdl: userHdl,
		bookHdl: bookHdl,
	}
}
