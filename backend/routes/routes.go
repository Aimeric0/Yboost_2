package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/aimeric0/yboost_2/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/recipes", controllers.GetAllRecipes)
		api.GET("/recipes/:id", controllers.GetRecipeByID)
		api.POST("/recipes", controllers.CreateRecipe)
		api.PUT("/recipes/:id", controllers.UpdateRecipe)
		api.DELETE("/recipes/:id", controllers.DeleteRecipe)
	}
}
