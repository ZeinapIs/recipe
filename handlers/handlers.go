// handlers/handlers.go
package handlers

import (
	"github.com/ZeinapIs/recipe/database"
	"github.com/gofiber/fiber/v2"
)

// AddRecipeHandler adds a new recipe to the database
func AddRecipeHandler(c *fiber.Ctx) error {
	var newRecipe models.Recipe

	// Parse request body to get recipe details
	if err := c.BodyParser(&newRecipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Perform database insertion
	if err := database.DB.Create(&newRecipe).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add recipe"})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Recipe added successfully"})
}

// UpdateRecipeHandler updates an existing recipe in the database
func UpdateRecipeHandler(c *fiber.Ctx) error {
	// Parse recipe ID from request parameters
	recipeID := c.Params("id")

	// Check if the recipe exists
	var existingRecipe models.Recipe
	if err := database.DB.First(&existingRecipe, recipeID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Parse request body to get updated recipe details
	var updatedRecipe models.Recipe
	if err := c.BodyParser(&updatedRecipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Update existing recipe
	database.DB.Model(&existingRecipe).Updates(&updatedRecipe)

	// Return success response
	return c.JSON(fiber.Map{"message": "Recipe updated successfully"})
}

// DeleteRecipeHandler deletes a recipe from the database
func DeleteRecipeHandler(c *fiber.Ctx) error {
	// Parse recipe ID from request parameters
	recipeID := c.Params("id")

	// Check if the recipe exists
	var existingRecipe models.Recipe
	if err := database.DB.First(&existingRecipe, recipeID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Delete the recipe
	database.DB.Delete(&existingRecipe)

	// Return success response
	return c.JSON(fiber.Map{"message": "Recipe deleted successfully"})
}

// GetRecipeByIDHandler retrieves a specific recipe by ID
func GetRecipeByIDHandler(c *fiber.Ctx) error {
	// Parse recipe ID from request parameters
	recipeID := c.Params("id")

	// Check if the recipe exists
	var recipe models.Recipe
	if err := database.DB.First(&recipe, recipeID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Return the recipe
	return c.JSON(recipe)
}

// GetAllRecipesHandler retrieves a list of all recipes
func GetAllRecipesHandler(c *fiber.Ctx) error {
	var recipes []models.Recipe

	// Fetch all recipes from the database
	database.DB.Find(&recipes)

	// Return the list of recipes
	return c.JSON(recipes)
}

// SearchRecipesByNameHandler searches for recipes by name
func SearchRecipesByNameHandler(c *fiber.Ctx) error {
	// Parse query parameter for the recipe name
	recipeName := c.Query("name")

	// Search for recipes by name
	var recipes []models.Recipe
	database.DB.Where("name LIKE ?", "%"+recipeName+"%").Find(&recipes)

	// Return the list of matching recipes
	return c.JSON(recipes)
}

// SearchRecipesByIngredientHandler searches for recipes by ingredient
func SearchRecipesByIngredientHandler(c *fiber.Ctx) error {
	// Parse query parameter for the ingredient
	ingredient := c.Query("ingredient")

	// Search for recipes by ingredient
	var recipes []models.Recipe
	database.DB.Where("ingredients LIKE ?", "%"+ingredient+"%").Find(&recipes)

	// Return the list of matching recipes
	return c.JSON(recipes)
}

// SearchRecipesByCategoryHandler searches for recipes by category (example)
func SearchRecipesByCategoryHandler(c *fiber.Ctx) error {
	// Parse query parameter for the category
	category := c.Query("category")
	// Search for recipes by category (assuming a 'category' field in the Recipe model)
	var recipes []models.Recipe
	database.DB.Where("category = ?", category).Find(&recipes)

	// Return the list of matching recipes
	return c.JSON(recipes)
}

// MarkAsLikedHandler handles marking a recipe as liked
func MarkAsLikedHandler(c *fiber.Ctx) error {
	// Parse recipe ID from request parameters
	recipeID := c.Params("id")

	// Check if the recipe exists
	var existingRecipe models.Recipe
	if err := database.DB.First(&existingRecipe, recipeID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Increment the 'Likes' count (assuming a 'Likes' field in the Recipe model)
	database.DB.Model(&existingRecipe).UpdateColumn("likes", existingRecipe.Likes+1)

	// Return success response
	return c.JSON(fiber.Map{"message": "Recipe marked as liked"})
}

// MarkAsDislikedHandler handles marking a recipe as disliked
func MarkAsDislikedHandler(c *fiber.Ctx) error {
	// Parse recipe ID from request parameters
	recipeID := c.Params("id")

	// Check if the recipe exists
	var existingRecipe models.Recipe
	if err := database.DB.First(&existingRecipe, recipeID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Increment the 'Dislikes' count (assuming a 'Dislikes' field in the Recipe model)
	database.DB.Model(&existingRecipe).UpdateColumn("dislikes", existingRecipe.Dislikes+1)

	// Return success response
	return c.JSON(fiber.Map{"message": "Recipe marked as disliked"})
}

// GetAllLikedRecipesHandler retrieves a list of all liked recipes
func GetAllLikedRecipesHandler(c *fiber.Ctx) error {
	var likedRecipes []models.Recipe

	// Fetch all recipes with at least one like
	database.DB.Where("likes > 0").Find(&likedRecipes)

	// Return the list of liked recipes
	return c.JSON(likedRecipes)
}

// GetAllDislikedRecipesHandler retrieves a list of all disliked recipes
func GetAllDislikedRecipesHandler(c *fiber.Ctx) error {
	var dislikedRecipes []models.Recipe

	// Fetch all recipes with at least one dislike
	database.DB.Where("dislikes > 0").Find(&dislikedRecipes)

	// Return the list of disliked recipes
	return c.JSON(dislikedRecipes)
}

// ErrorHandler is a middleware for handling errors
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Log the error or perform additional error handling
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
}
