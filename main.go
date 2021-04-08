package main

import (	
"github.com/chaitanyav98/golang-fiber/book"
"github.com/gofiber/fiber"
"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}


func main(){
	app:= fiber.New()

	setupRoutes(app);
	app.Listen(4000)
	defer database.DBConn.Close()

}

//

