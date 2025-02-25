package controllers

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

var books []models.Book

func AddBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	books = append(books, book)
	return c.JSON(fiber.Map{"message": "Book added"})
}

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
