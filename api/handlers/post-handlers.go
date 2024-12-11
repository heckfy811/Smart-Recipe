package handlers

import (
	"SmartRecipe/database"
	"SmartRecipe/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostRecipeHandler(c *gin.Context) {
	var req models.Recipe
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.Database.Recipes.AddRecipe(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func PostIngredientsHandler(c *gin.Context) {
	var req []models.Ingredient
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for id, r := range req {
		r.Id = id + 1
		err := database.Database.Ingredients.AddIngredient(&r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
