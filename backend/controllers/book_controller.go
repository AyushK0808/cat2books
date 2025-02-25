package controllers

import (
	"backend/config"
	"backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

// AddBook - Controller to add a book
func AddBook(c *fiber.Ctx) error {
	// Parse JSON request
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate input
	if book.Name == "" || book.Subject == "" || book.Code == "" || len(book.Slots) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Get Firebase Firestore client
	client, err := config.FirebaseFirestore()
	if err != nil {
		log.Println("Firestore error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Add book to Firestore
	docRef, _, err := client.Collection("books").Add(context.Background(), book)
	if err != nil {
		log.Println("Error adding book:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add book"})
	}

	book.ID = docRef.ID // Assign Firestore document ID

	return c.JSON(fiber.Map{
		"message": "Book added successfully",
		"book":    book,
	})
}

// GetBooks - Retrieve all books
func GetBooks(c *fiber.Ctx) error {
	client, err := config.FirebaseFirestore()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer client.Close()

	// Fetch books from Firestore
	var books []models.Book
	iter := client.Collection("books").Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching books"})
		}

		var book models.Book
		doc.DataTo(&book)
		book.ID = doc.Ref.ID
		books = append(books, book)
	}

	return c.JSON(books)
}
