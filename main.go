package main

import (
	"github.com/azizanhakim/notes-api-fiber/database"
	"github.com/azizanhakim/notes-api-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
