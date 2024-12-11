package handlers

import (
	"SmartRecipe/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRecipeHandler(c *gin.Context) {
	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.Database.Recipes.DeleteRecipe(recipeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
}

func DeleteIngredientHandler(c *gin.Context) {

	IngredientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if IngredientId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingredient id should be greater than 0"})
		return
	}
	err = database.Database.Ingredients.DeleteIngredient(IngredientId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ingredient deleted"})
}
