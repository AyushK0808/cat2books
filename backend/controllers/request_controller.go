package controllers

import (
	"backend/config"
	"backend/models"
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

// CreateBookRequest - User requests a book
func CreateRequest(c *fiber.Ctx) error {
	var request models.Request

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	// Validate input
	if request.BookID == "" || request.Requester == "" || request.Owner == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Set default values
	request.Status = "pending"
	request.CreatedAt = time.Now()

	client, err := config.FirebaseFirestore()
	if err != nil {
		log.Println("Firestore error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Store book request
	ctx := context.Background()
	docRef, _, err := client.Collection("book_requests").Add(ctx, request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create request"})
	}

	request.ID = docRef.ID

	return c.JSON(fiber.Map{
		"message": "Book request sent successfully",
		"request": request,
	})
}

// GetUserRequests - Retrieve all requests made by a user
func GetUserRequests(c *fiber.Ctx) error {
	userID := c.Params("userID") // User who made the requests

	client, err := config.FirebaseFirestore()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	var requests []models.Request
	iter := client.Collection("book_requests").Where("Requester", "==", userID).Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching requests"})
		}

		var request models.Request
		doc.DataTo(&request)
		request.ID = doc.Ref.ID
		requests = append(requests, request)
	}

	return c.JSON(requests)
}

// GetOwnerRequests - Retrieve all requests for books owned by a user
func GetOwnerRequests(c *fiber.Ctx) error {
	ownerID := c.Params("ownerID") // Book owner

	client, err := config.FirebaseFirestore()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	var requests []models.Request
	iter := client.Collection("book_requests").Where("Owner", "==", ownerID).Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching requests"})
		}

		var request models.Request
		doc.DataTo(&request)
		request.ID = doc.Ref.ID
		requests = append(requests, request)
	}

	return c.JSON(requests)
}

// UpdateRequestStatus - Accept or reject a request
func UpdateRequestStatus(c *fiber.Ctx) error {
	requestID := c.Params("id")
	var updateData struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if updateData.Status != "accepted" && updateData.Status != "rejected" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status"})
	}

	client, err := config.FirebaseFirestore()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Update request status
	ctx := context.Background()
	_, err = client.Collection("book_requests").Doc(requestID).Update(ctx, []firestore.Update{
		{Path: "Status", Value: updateData.Status},
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update status"})
	}

	return c.JSON(fiber.Map{"message": "Request status updated successfully"})
}
