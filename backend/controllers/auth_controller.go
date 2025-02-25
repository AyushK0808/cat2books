package controllers

import (
	"os"
	"time"

	"backend/config"
	"context"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	client, err := config.FirebaseApp.Auth(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Firebase error"})
	}

	user, err := client.GetUserByEmail(c.Context(), request.Email)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.UID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "JWT error"})
	}

	return c.JSON(fiber.Map{"token": signedToken, "user": user.Email})
}

func Signup(c *fiber.Ctx) error {
	var request SignupRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	client, err := config.FirebaseApp.Auth(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Firebase error"})
	}

	// Create a new user in Firebase Auth
	params := (&auth.UserToCreate{}).
		Email(request.Email).
		Password(request.Password)

	userRecord, err := client.CreateUser(context.Background(), params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error creating user", "details": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"userId":  userRecord.UID,
	})
}
