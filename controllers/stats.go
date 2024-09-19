package controllers

import (
	"finance-tracker/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatController struct {
	statService *services.StatService
}

func NewStatController(statService *services.StatService) *StatController {
	return &StatController{statService}
}

func (c *StatController) GetStats(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	stat, err := c.statService.GetStats(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": stat})
}

func (c *StatController) GetYearlyStats(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	year := ctx.Query("year")
	if year == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Year parameter is required"})
		return
	}

	//covert year from string to int
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid year format"})
		return
	}

	stat, err := c.statService.GetYearlyStats(userID.(uint), yearInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": stat})
}
