package route

import (
	"github.com/MrJSdev/goBookStore/api"
	"github.com/kataras/iris/v12"
)

func RegisterBookRoutes(app *iris.Application) {
	app.Get("/book", api.GetBooks)
	app.Get("/book/first", api.GetFirstBook)
	app.Post("/book/create", api.AddBook)
	app.Delete("/book/{id}", api.DeleteBook)
	app.Get("/book/{id}", api.GetBookByID)
	app.Put("/book/{id}", api.UpdateBook)
}
