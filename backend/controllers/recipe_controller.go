package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/aimeric0/yboost_2/models"
)

var recipes = []models.Recipe{
	{ID: 1, Name: "Pasta Carbonara", Country: "Italy", Ingredients: "Pasta, eggs, cheese, bacon", Type: "main", IsVegan: false},
	{ID: 2, Name: "Sushi", Country: "Japan", Ingredients: "Rice, fish, seaweed", Type: "main", IsVegan: false},
}

func GetAllRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func GetRecipeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, recipe := range recipes {
		if int(recipe.ID) == id {
			c.JSON(http.StatusOK, recipe)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
}

func CreateRecipe(c *gin.Context) {
	var newRecipe models.Recipe
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRecipe.ID = uint(len(recipes) + 1)
	recipes = append(recipes, newRecipe)
	c.JSON(http.StatusCreated, newRecipe)
}

func UpdateRecipe(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updated models.Recipe
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, recipe := range recipes {
		if int(recipe.ID) == id {
			updated.ID = recipe.ID
			recipes[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
}

func DeleteRecipe(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, recipe := range recipes {
		if int(recipe.ID) == id {
			recipes = append(recipes[:i], recipes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
}
