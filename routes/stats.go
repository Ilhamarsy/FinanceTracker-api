package routes

import (
	"finance-tracker/controllers"
	"finance-tracker/middlewares"

	"github.com/gin-gonic/gin"
)

func StatRoutes(router *gin.Engine, statController *controllers.StatController) {
	api := router.Group("/api")
	{
		protected := api.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		protected.GET("/stats", statController.GetStats)
		protected.GET("/stats-yearly", statController.GetYearlyStats)
	}
}
