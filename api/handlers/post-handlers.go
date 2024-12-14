package handlers

import (
	"SmartRecipe/database"
	"SmartRecipe/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostIngredientsHandler(c *gin.Context) {
	var req []*models.Ingredient
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, r := range req {
		err := database.Database.Ingredients.AddIngredient(r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func PostRecipesHandler(c *gin.Context) {
	type FullRecipe struct {
		Title         string              `json:"title"`
		Description   string              `json:"description"`
		Instructions  string              `json:"instructions"`
		Kilocalories  float64             `json:"kilocalories"`
		Proteins      float64             `json:"proteins"`
		Fats          float64             `json:"fats"`
		Carbohydrates float64             `json:"carbohydrates"`
		Score         float64             `json:"score"`
		Ingredients   []models.Ingredient `json:"ingredients"`
	}

	var recipes []FullRecipe
	if err := c.ShouldBindJSON(&recipes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range recipes {
		recipe := &models.Recipe{
			Title:         req.Title,
			Description:   req.Description,
			Instructions:  req.Instructions,
			Kilocalories:  req.Kilocalories,
			Proteins:      req.Proteins,
			Fats:          req.Fats,
			Carbohydrates: req.Carbohydrates,
			Score:         req.Score,
		}

		id, err := database.Database.Recipes.AddRecipe(recipe)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, ingredient := range req.Ingredients {
			ingredient.RecipeId = id
			err = database.Database.Ingredients.AddIngredient(&ingredient)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
