package handlers

import (
	"net/http"

	"git-practice-gin/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var list []models.Category
	for _, cat := range categories {
		list = append(list, cat)
	}
	c.JSON(http.StatusOK, list)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name required"})
		return
	}
	category.ID = nextCategoryID
	nextCategoryID++
	categories[category.ID] = category
	c.JSON(http.StatusCreated, category)
}
