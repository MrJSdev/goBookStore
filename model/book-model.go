package model

import (
	"fmt"

	"github.com/MrJSdev/goBookStore/config"
	"github.com/MrJSdev/goBookStore/entity"
	"gorm.io/gorm"
)

func SaveBook(book *entity.Book) {
	var db = config.GetDB()
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
	var db = config.GetDB()
	var book entity.Book

	return book, db.Model(entity.Book{}).Where("ID=?", bookId).Find(&book)
}

func GetFirstBook(book *entity.Book) {
	book.Title = "title from backend"
}

func GetBooksList() []entity.Book {
	var db = config.GetDB()

	var books []entity.Book
	db.Find(&books)
	return books
}

func DeleteBook(bookId uint) entity.Book {
	foundBook, findDb := GetBookByID(bookId)

	findDb.Delete(&foundBook)

	return foundBook
}
