package service

import (
	"github.com/MrJSdev/goBookStore/entity"
	"github.com/MrJSdev/goBookStore/model"
)

// Business logic services
func SaveBook(book *entity.Book) {

	model.SaveBook(book)
}

func GetBookByID(bookId uint) entity.Book {
	book, _ := model.GetBookByID(bookId)

	return book
}

func UpdateBook(bookId uint, book *entity.Book) {

	model.UpdateBook(bookId, book)
}

func DeleteBook(bookId uint) entity.Book {

	return model.DeleteBook(bookId)
}

func GetFirstBook(book *entity.Book) {

	model.GetFirstBook(book)
}

func GetBooks() []entity.Book {

	return model.GetBooksList()
}
