package model

import (
	"time"

	"gorm.io/gorm"
)

// https://gorm.io/docs/models.html
// bagaimana caranya kita menyambungkan
// golang struct dengan gorm
type Book struct {
	// - di json, menandakan golang akan mengabaikan properti
	// - di gorm, menandakan gorm akan mengabaikan properti
	Id        		uint64         `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	NameBook      string         `json:"name_book" gorm:"column:name_book;not null"`
	Author     		string         `json:"author" gorm:"column:author;not null"`
	Delete    		bool           `json:"-" gorm:"-"`
	CreatedAt 		time.Time      `json:"created_at"`
	UpdatedAt 		time.Time      `json:"updated_at"`
	DeletedAt 		gorm.DeletedAt `json:"-" gorm:"index"`
	// gorm.Model
}