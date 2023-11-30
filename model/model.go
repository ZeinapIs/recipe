// models/recipe.go
package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
	Likes        int    `json:"likes"`
	Dislikes     int    `json:"dislikes"`
}
