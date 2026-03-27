package handlers

import (
	"net/http"

	"git-practice-gin/models"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var list []models.Author
	for _, a := range authors {
		list = append(list, a)
	}
	c.JSON(http.StatusOK, list)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if author.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name required"})
		return
	}
	author.ID = nextAuthorID
	nextAuthorID++
	authors[author.ID] = author
	c.JSON(http.StatusCreated, author)
}
