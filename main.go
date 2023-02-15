package main

import (
	"github.com/MrJSdev/goBookStore/config"
	"github.com/MrJSdev/goBookStore/route"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	config.InitConnection()

	// Test route
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"test": "Welcome riyaz",
		})
	})

	// Book routes
	route.RegisterBookRoutes(app)

	app.Listen(":8080")
}
