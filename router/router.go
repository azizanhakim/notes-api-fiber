package router

import (
	noteRoutes "github.com/azizanhakim/notes-api-fiber/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// then we must define our routes function
	noteRoutes.SetupNoteRoutes(app)

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(api)
}
