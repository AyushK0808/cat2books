package controllers

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

var requests []models.Request

func RequestBook(c *fiber.Ctx) error {
	var req models.Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	req.Status = "Pending"
	requests = append(requests, req)
	return c.JSON(fiber.Map{"message": "Request sent"})
}
