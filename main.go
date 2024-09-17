package main

import (
	"finance-tracker/config"
	"finance-tracker/controllers"
	"finance-tracker/repositories"
	"finance-tracker/routes"
	"finance-tracker/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Database connection
	database := config.NewDatabase(cfg)
	database.ConnectDB()

	// Initialize repositories and services
	authRepo := repositories.NewAuthRepository(config.DB)
	authService := services.NewAuthService(cfg, authRepo)
	authController := controllers.NewAuthController(authService)

	categoryRepo := repositories.NewCategoryRepository(config.DB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	incomeRepo := repositories.NewIncomeRepository(config.DB)
	incomeService := services.NewIncomeService(incomeRepo)
	incomeController := controllers.NewIncomeController(incomeService)

	expenseRepo := repositories.NewExpenseRepository(config.DB)
	expenseService := services.NewExpenseService(expenseRepo, categoryRepo)
	expenseController := controllers.NewExpenseController(expenseService)

	statService := services.NewStatService(expenseRepo, incomeRepo)
	statController := controllers.NewStatController(statService)

	// Set up routes
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "UPDATE", "DELETE"}, // Allowed methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.CategoryRoutes(router, categoryController)
	routes.AuthRoutes(router, authController)
	routes.IncomeRoutes(router, incomeController)
	routes.ExpenseRoutes(router, expenseController)
	routes.StatRoutes(router, statController)

	port := cfg.Port
	if port == "0" {
		port = "8080" // Default port if not defined
	}
	router.Run(":" + port)
}
