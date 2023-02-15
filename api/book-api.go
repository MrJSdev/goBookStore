package api

import (
	"github.com/MrJSdev/goBookStore/entity"
	"github.com/MrJSdev/goBookStore/service"
	"github.com/kataras/iris/v12"
)

func AddBook(ctx iris.Context) {
	var book entity.Book

	ctx.ReadBody(&book)

	service.SaveBook(&book)

	ctx.JSON(book)
}

func UpdateBook(ctx iris.Context) {
	var book entity.Book
	var bookId, err = ctx.Params().GetUint("id")

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("There is something wrong with id param").DetailErr(err))
		return
	}

	ctx.ReadBody(&book)

	service.UpdateBook(bookId, &book)

	ctx.JSON(book)
}

func DeleteBook(ctx iris.Context) {
	var bookId, err = ctx.Params().GetUint("id")

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("There is something wrong with id param").DetailErr(err))
		return
	}

	var book = service.DeleteBook(bookId)

	if book.ID == 0 {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("No records found"))
	}
}

func GetBookByID(ctx iris.Context) {
	var bookId, err = ctx.Params().GetUint("id")

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.JSON(service.GetBookByID(bookId))
}

func GetFirstBook(ctx iris.Context) {
	var book entity.Book

	service.GetFirstBook(&book)

	ctx.JSON(book)
}

func GetBooks(ctx iris.Context) {

	ctx.JSON(service.GetBooks())
}
