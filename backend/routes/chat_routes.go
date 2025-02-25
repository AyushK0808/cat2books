package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupChatRoutes - Define chat routes
func SetupChatRoutes(app *fiber.App) {
	chat := app.Group("/chat")

	chat.Post("/start", controllers.StartChat)           // Start a new chat
	chat.Get("/:chat_id", controllers.GetChat)           // Fetch chat messages
	chat.Post("/:chat_id/send", controllers.SendMessage) // Send a message
}
