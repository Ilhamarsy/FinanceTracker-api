package routes

import (
	"finance-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	api := router.Group("/api")
	{
		// api.POST("/category", middlewares.JWTAuthMiddleware(), categoryController.CreateCategory)
		// api.GET("/categories", middlewares.JWTAuthMiddleware(), categoryController.GetCategories)
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
	}
}
