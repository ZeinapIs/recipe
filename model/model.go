// models/recipe.go
package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	ID           int
	Name         string
	Ingredients  string
	Instructions string
	Likes        uint
	Dislikes     uint
	// Add other fields as needed
}
