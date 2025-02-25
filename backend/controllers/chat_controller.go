package controllers

import (
	"backend/config"
	"backend/models"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// StartChat - Initializes a chat between two users
func StartChat(c *fiber.Ctx) error {
	var request struct {
		RequestID string `json:"request_id"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	client, err := config.FirebaseFirestore()
	if err != nil {
		log.Println("Firestore error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer client.Close()

	// Fetch the book request details
	doc, err := client.Collection("book_requests").Doc(request.RequestID).Get(context.Background())
	if err != nil || !doc.Exists() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Request not found"})
	}

	var bookRequest models.Request
	doc.DataTo(&bookRequest)

	// Generate unique chat ID
	chatID := bookRequest.Requester + "_" + bookRequest.Owner

	// Check if chat already exists
	chatDoc, err := client.Collection("chats").Doc(chatID).Get(context.Background())
	if err == nil && chatDoc.Exists() {
		return c.JSON(fiber.Map{"message": "Chat already exists", "chat_id": chatID, "url": "/chat/" + chatID})
	}

	// Create a new chat
	chat := models.Chat{
		ID:        chatID,
		BookID:    bookRequest.BookID,
		Requester: bookRequest.Requester,
		Owner:     bookRequest.Owner,
		Messages:  []models.ChatMessage{},
	}

	_, err = client.Collection("chats").Doc(chatID).Set(context.Background(), chat)
	if err != nil {
		log.Println("Error creating chat:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create chat"})
	}

	return c.JSON(fiber.Map{
		"message": "Chat initialized successfully",
		"chat_id": chatID,
		"url":     "/chat/" + chatID,
	})
}

// GetChat - Fetches chat messages
func GetChat(c *fiber.Ctx) error {
	chatID := c.Params("chat_id")

	client, err := config.FirebaseFirestore()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer client.Close()

	// Fetch chat document
	doc, err := client.Collection("chats").Doc(chatID).Get(context.Background())
	if err != nil || !doc.Exists() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chat not found"})
	}

	var chat models.Chat
	doc.DataTo(&chat)

	return c.JSON(chat)
}

// SendMessage - Sends a message in a chat
func SendMessage(c *fiber.Ctx) error {
	chatID := c.Params("chat_id")

	var msg models.ChatMessage
	if err := c.BodyParser(&msg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid message format"})
	}

	// Add timestamp
	msg.Timestamp = time.Now().Format(time.RFC3339)

	client, err := config.FirebaseFirestore()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer client.Close()

	// Fetch chat
	docRef := client.Collection("chats").Doc(chatID)
	doc, err := docRef.Get(context.Background())
	if err != nil || !doc.Exists() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chat not found"})
	}

	var chat models.Chat
	doc.DataTo(&chat)

	// Append new message
	chat.Messages = append(chat.Messages, msg)

	_, err = docRef.Set(context.Background(), chat)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send message"})
	}

	return c.JSON(fiber.Map{"message": "Message sent successfully"})
}
