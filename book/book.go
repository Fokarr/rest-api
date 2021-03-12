package book

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"awesomeProject/database"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.ID == 0 {
		return c.Status(404).SendString("No User with the given ID was found.")
	}
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book with the given ID was found.")
	}
	titlePointer := &book.Title
	title := *titlePointer
	db.Delete(&book, id)
	return c.Status(200).SendString("The Book:" + title + " was deleted.")
}

