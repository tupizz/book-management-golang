package models

import (
	"book-management-system-golang/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Books struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Books) Create() *Books {
	db.Create(&b)
	return b
}

func GetAll() []Books {
	var books []Books
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Books, *gorm.DB) {
	var book Books
	d := db.Where("ID=?", id).Find(&book)
	return &book, d
}

func DeleteBook(ID int64) Books {
	var book Books
	db.Where("ID=?", ID).Delete(book)
	return book
}

func init() {
	config.Connect()
	db = config.GetDatabase()
	db.AutoMigrate(&Books{})
}
