package server

import (
	"library-api/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type postData struct {
	Name   string `json:"name" form:"name"`
	Author string `json:"author" form:"author"`
	Id 	   int `json:"id" form:"id"`
}

func RunServer() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/search/", getFilteredBooks)
	app.Get("/all/", getAllBooks)
	app.Post("/book/add/", addBook)
	app.Delete("/book/:id/", deleteBook)
	app.Patch("/book/:id/", editBook)


	app.Listen(":8080")
}

func getFilteredBooks(c *fiber.Ctx) error {
	rData := new(postData)
	if err := c.BodyParser(rData); err != nil {
		return err
	}
	books := db.GetAllFilteredBooks(rData.Name, rData.Author)
	return c.JSON(books)
}

func getAllBooks(c *fiber.Ctx) error {
	books := db.GetAllBooks()
	return c.JSON(books)
}

func addBook(c *fiber.Ctx) error {
	rData := new(postData)
	if err := c.BodyParser(rData); err != nil {
		return err
	}
	if db.AddBook(rData.Name, rData.Author) {
		return c.SendStatus(200)
	}
	return c.SendStatus(400)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := c.ParamsInt("id")
	if bookId != 0 && err == nil {
		if db.DeleteBook(bookId) {
			return c.SendStatus(200)
		}
		return c.SendStatus(500)
	}
	return c.SendStatus(400)
}

func editBook(c *fiber.Ctx) error {
	bookId, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}
	
	return c.SendStatus(200)
}