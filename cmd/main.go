// main.go
package main

import (
	"github.com/ZeinapIs/recipe/database"
	"github.com/ZeinapIs/recipe/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
