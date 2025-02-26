package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // Allow requests from your frontend
		AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,PATCH,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Register Routes
	routes.SetupAuthRoutes(app)
	routes.SetupBookRoutes(app)
	routes.SetupBookRequestRoutes(app)
	routes.SetupChatRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // Changed to 4000 since backend runs on this port
	}

	for _, route := range app.GetRoutes() {
		fmt.Println(route.Method, route.Path)
	}

	fmt.Println("Server running on port " + port)
	log.Fatal(app.Listen(":" + port))
}
