package main

import (	
"github.com/chaitanyav98/go-fiber/book"
"github.com/gofiber/fiber"
)
//import ("goland-crud/book")

func helloWorld(c *fiber.Ctx){
	c.Send("hello world")
}
func setupRoutes(app *fiber.App){
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}
func main(){
	app:= fiber.New()

	setupRoutes(app);
	app.Listen(4000)

}

//

