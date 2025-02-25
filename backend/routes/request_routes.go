package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupBookRequestRoutes - Routes for book requests
func SetupBookRequestRoutes(app *fiber.App) {
	api := app.Group("/requests")

	api.Post("/add", controllers.CreateRequest)
	api.Get("/user/:userID", controllers.GetUserRequests)    // Requests made by a user
	api.Get("/owner/:ownerID", controllers.GetOwnerRequests) // Requests for a user's books
	api.Patch("/:id", controllers.UpdateRequestStatus)       // Accept/reject request
}
