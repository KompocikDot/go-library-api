package server

import (
	"library-api/pkg/db"
	"library-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunServer() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/search/", getFilteredBooks)
	app.Get("/all/", getAllBooks)
	app.Post("/book/", addBook)
	app.Patch("/book/:id/", patchBook)
	app.Delete("/book/:id/", deleteBook)

	app.Listen(":8080")
}

func getFilteredBooks(c *fiber.Ctx) error {
	rData := new(utils.ReqData)
	if err := c.BodyParser(rData); err != nil {
		return err
	}
	books := db.GetAllFilteredBooks(*rData.Name, *rData.Author)
	return c.JSON(books)
}

func getAllBooks(c *fiber.Ctx) error {
	books := db.GetAllBooks()
	return c.JSON(books)
}

func addBook(c *fiber.Ctx) error {
	rData := new(utils.ReqData)
	if err := c.BodyParser(rData); err != nil {
		return err
	}
	if db.AddBook(*rData.Name, *rData.Author) {
		return c.SendStatus(fiber.StatusOK)
	}
	return c.SendStatus(fiber.StatusBadRequest)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := c.ParamsInt("id")
	if bookId != 0 && err == nil {
		if db.DeleteBook(bookId) {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusBadRequest)
}

func patchBook(c *fiber.Ctx) error {
	rData := new(utils.ReqData)
	if err := c.BodyParser(rData); err != nil {
		return err
	}

	bookId, err := c.ParamsInt("id")
	err = db.EditBook(bookId, *rData)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}
