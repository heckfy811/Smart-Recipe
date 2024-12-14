package handlers

import (
	"SmartRecipe/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRecipeHandler(c *gin.Context) {
	var req struct {
		Id int `json:"id"`
	}
	err := c.ShouldBindJSON(&req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.Database.Recipes.DeleteRecipe(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
}

func DeleteIngredientHandler(c *gin.Context) {
	var req struct {
		Id int `json:"id"`
	}
	err := c.ShouldBindJSON(&req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.Database.Ingredients.DeleteIngredient(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ingredient deleted"})
}
