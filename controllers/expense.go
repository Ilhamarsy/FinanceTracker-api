package controllers

import (
	"finance-tracker/models"
	"finance-tracker/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	expenseService *services.ExpenseService
}

func NewExpenseController(expenseService *services.ExpenseService) *ExpenseController {
	return &ExpenseController{expenseService}
}

func (e *ExpenseController) AddExpense(c *gin.Context) {
	var req models.ExpenseRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	req.UserID = userID.(uint)

	err := e.expenseService.AddExpense(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Expense added successfully"})
}

func (e *ExpenseController) GetExpenses(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	expenses, err := e.expenseService.GetExpenses(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": expenses})
}

func (e *ExpenseController) UpdateExpense(ctx *gin.Context) {
	var req models.ExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	req.UserID = userID.(uint)

	if expenseID, err := strconv.ParseUint(ctx.Param("id"), 10, 64); err == nil {
		req.Id = uint(expenseID)
		if err := e.expenseService.UpdateExpense(&req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid id"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Expense updated successfully"})
}

func (e *ExpenseController) DeleteExpense(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	if expenseID, err := strconv.ParseUint(ctx.Param("id"), 10, 64); err == nil {
		if err := e.expenseService.DeleteExpense(uint(expenseID), userID.(uint)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid id"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Expense deleted successfully"})
}
