package controllers

import (
	"finance-tracker/models"
	"finance-tracker/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IncomeController struct {
	incomeService *services.IncomeService
}

func NewIncomeController(incomeService *services.IncomeService) *IncomeController {
	return &IncomeController{incomeService}
}

func (i *IncomeController) AddIncome(ctx *gin.Context) {
	var income models.Income
	if err := ctx.ShouldBindJSON(&income); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	income.UserID = userID.(uint)

	if err := i.incomeService.AddIncome(&income); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Income added successfully"})
}

func (i *IncomeController) GetIncomes(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	incomes, err := i.incomeService.GetIncomes(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": incomes})
}

func (i *IncomeController) UpdateIncome(ctx *gin.Context) {
	var income models.Income
	if err := ctx.ShouldBindJSON(&income); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	income.UserID = userID.(uint)

	if incomeID, err := strconv.ParseUint(ctx.Param("id"), 10, 64); err == nil {
		income.ID = uint(incomeID)
		if err := i.incomeService.UpdateIncome(&income); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid id"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Income updated successfully"})
}

func (i *IncomeController) DeleteIncome(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	if incomeID, err := strconv.ParseUint(ctx.Param("id"), 10, 64); err == nil {
		if err := i.incomeService.DeleteIncome(uint(incomeID), userID.(uint)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid id"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Income deleted successfully"})
}
