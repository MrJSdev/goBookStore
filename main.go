package main

import (
	"github.com/MrJSdev/goBookStore/api"
	"github.com/MrJSdev/goBookStore/model"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	model.InitDbConnection()

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"test": "Welcome riyaz",
		})
	})

	app.Get("/book", api.GetBooks)
	app.Get("/book/first", api.GetFirstBook)
	app.Post("/book/create", api.AddBook)
	app.Delete("/book/{id}", api.DeleteBook)
	app.Get("/book/{id}", api.GetBookByID)
	app.Put("/book/{id}", api.UpdateBook)

	app.Listen(":8080")
}
