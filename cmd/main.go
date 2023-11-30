// main.go
package main

import (
	"github.com/ZeinapIs/recipe/database"
	"github.com/ZeinapIs/recipe/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup database
	database.InitDatabase()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":8080")
}
