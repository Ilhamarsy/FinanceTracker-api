package routes

import (
	"finance-tracker/controllers"
	"finance-tracker/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, categoryController *controllers.CategoryController) {
	api := router.Group("/api")
	{
		protected := api.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		protected.POST("/category", categoryController.CreateCategory)
		protected.GET("/categories", categoryController.GetCategories)
	}
}
