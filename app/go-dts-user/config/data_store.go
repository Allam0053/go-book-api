package config

import "github.com/Calmantara/go-dts-user/module/model"

type DataStore struct {
	UserData map[uint64]model.User
	BookData map[uint64]model.Book
}

func ConnectDataStore() (ds DataStore) {
	// init map
	userData := make(map[uint64]model.User)
	bookData := make(map[uint64]model.Book)

	return DataStore{
		UserData: userData,
		BookData: bookData,
	}
}
