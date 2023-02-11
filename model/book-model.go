package model

import (
	"fmt"

	"github.com/MrJSdev/goBookStore/config"
	"github.com/MrJSdev/goBookStore/entity"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDbConnection() {
	dbValue, err := config.GetDB()

	db = dbValue

	db.AutoMigrate(&entity.Book{})

	if err != nil {
		fmt.Printf("there was something wrong to connect DB: %s", err.Error())
	}

}

func SaveBook(book *entity.Book) {
	db.AutoMigrate(&entity.Book{})

	db.Create(&book)
}

func UpdateBook(bookId uint, book *entity.Book) {
	foundBook, findDb := GetBookByID(bookId)

	if book.Title != "" {
		foundBook.Title = book.Title
	}

	if book.Description != "" {
		foundBook.Description = book.Description
	}

	fmt.Println(foundBook, "foundbook")

	findDb.Save(&foundBook)

	book = &foundBook
	fmt.Println(book)
}

func GetBookByID(bookId uint) (entity.Book, *gorm.DB) {
	var book entity.Book

	return book, db.Model(entity.Book{}).Where("ID=?", bookId).Find(&book)
}

func GetFirstBook(book *entity.Book) {
	book.Title = "title from backend"
}

func GetBooksList() []entity.Book {
	var books []entity.Book
	db.Find(&books)
	return books
}

func DeleteBook(bookId uint) entity.Book {
	foundBook, findDb := GetBookByID(bookId)

	findDb.Delete(&foundBook)

	return foundBook
}
