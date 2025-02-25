package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(app *fiber.App) {
	bookGroup := app.Group("/books")
	bookGroup.Post("/", controllers.AddBook)
	bookGroup.Get("/", controllers.GetBooks)
}
