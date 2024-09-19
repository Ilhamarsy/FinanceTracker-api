package controllers

import (
	"finance-tracker/models"
	"finance-tracker/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{categoryService}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	category.UserID = userID.(uint)

	if err := c.categoryService.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Category created successfully"})
}

func (c *CategoryController) GetCategories(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	categories, err := c.categoryService.GetCategories(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Unable to fetch categories"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": categories})
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User ID not found"})
		return
	}

	if categoryID, err := strconv.ParseUint(ctx.Param("id"), 10, 64); err == nil {
		if err := c.categoryService.DeleteCategory(uint(categoryID), userID.(uint)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid id"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Category deleted successfully"})
}
