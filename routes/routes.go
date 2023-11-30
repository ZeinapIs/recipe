package routes

import (
	"github.com/ZeinapIs/recipe/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	// Add a Recipe
	app.Post("/recipes", handlers.AddRecipeHandler)

	// Update a Recipe
	app.Put("/recipes/:id", handlers.UpdateRecipeHandler)

	// Delete a Recipe
	app.Delete("/recipes/:id", handlers.DeleteRecipeHandler)

	// Get Recipe by ID
	app.Get("/recipes/:id", handlers.GetRecipeByIDHandler)

	// Get All Recipes
	app.Get("/recipes", handlers.GetAllRecipesHandler)

	// Search Recipes by Name
	app.Get("/recipes/search", handlers.SearchRecipesByNameHandler)

	// Search Recipes by Ingredient
	app.Get("/recipes/search", handlers.SearchRecipesByIngredientHandler)

	// Search Recipes by Category (Assuming a 'Category' Field in Recipe Model)
	app.Get("/recipes/search", handlers.SearchRecipesByCategoryHandler)

	// Mark as Liked
	app.Post("/recipes/:id/like", handlers.MarkAsLikedHandler)

	// Mark as Disliked
	app.Post("/recipes/:id/dislike", handlers.MarkAsDislikedHandler)

	// Get All Liked Recipes
	app.Get("/recipes/liked", handlers.GetAllLikedRecipesHandler)

	// Get All Disliked Recipes
	app.Get("/recipes/disliked", handlers.GetAllDislikedRecipesHandler)
}
