package routes

import (
	"finance-tracker/controllers"
	"finance-tracker/middlewares"

	"github.com/gin-gonic/gin"
)

func ExpenseRoutes(router *gin.Engine, expenseController *controllers.ExpenseController) {
	api := router.Group("/api")
	{
		protected := api.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		protected.POST("/expense", expenseController.AddExpense)
		protected.GET("/expenses", expenseController.GetExpenses)
		protected.PUT("/expense/:id", expenseController.UpdateExpense)
		protected.DELETE("/expense/:id", expenseController.DeleteExpense)
	}
}
