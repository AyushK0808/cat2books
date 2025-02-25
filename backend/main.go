package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.InitFirebase()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Println("Error:", err)
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		},
	})

	// Register Routes
	routes.SetupAuthRoutes(app)
	routes.SetupBookRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	for _, route := range app.GetRoutes() {
		fmt.Println(route.Method, route.Path)
	}

	fmt.Println("Server running on port " + port)
	log.Fatal(app.Listen(":" + port))
}
