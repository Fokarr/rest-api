package main

import (
	"awesomeProject/book"
	"awesomeProject/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("\"I made a ☕ for you!\"")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected.")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated.")
}

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", helloWorld)
	setupRoutes(app)

	// Database
	initDatabase()

	log.Fatal(app.Listen(":3000"))
}
