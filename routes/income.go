package routes

import (
	"finance-tracker/controllers"
	"finance-tracker/middlewares"

	"github.com/gin-gonic/gin"
)

func IncomeRoutes(router *gin.Engine, incomeController *controllers.IncomeController) {
	api := router.Group("/api")
	{
		protected := api.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		protected.POST("/income", incomeController.AddIncome)
		protected.GET("/incomes", incomeController.GetIncomes)
		protected.PUT("/income/:id", incomeController.UpdateIncome)
		protected.DELETE("/income/:id", incomeController.DeleteIncome)
	}
}
