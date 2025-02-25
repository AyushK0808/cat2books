package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/signup", controllers.Signup)
	authGroup.Post("/login", controllers.Login)
}
