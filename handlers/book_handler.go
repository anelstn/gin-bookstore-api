package handlers

import (
	"net/http"
	"sort"
	"strconv"

	"git-practice-gin/models"

	"github.com/gin-gonic/gin"
)

var books = make(map[int]models.Book)
var authors = make(map[int]models.Author)
var categories = make(map[int]models.Category)

var nextBookID = 1
var nextAuthorID = 1
var nextCategoryID = 1

func GetBooks(c *gin.Context) {
	categoryFilter := c.Query("category")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	limit := 10
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
		limit = l
	}
	var filtered []models.Book
	for _, b := range books {
		if categoryFilter != "" {
			if cat, exists := categories[b.CategoryID]; !exists || cat.Name != categoryFilter {
				continue
			}
		}
		filtered = append(filtered, b)
	}
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].ID < filtered[j].ID
	})
	start := (page - 1) * limit
	end := start + limit
	if start >= len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}
	c.JSON(http.StatusOK, filtered[start:end])
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if book.Title == "" || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and price must be > 0"})
		return
	}
	book.ID = nextBookID
	nextBookID++
	books[book.ID] = book
	c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	book, exists := books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	var updated models.Book
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if updated.Title == "" || updated.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and price must be > 0"})
		return
	}
	updated.ID = id
	books[id] = updated
	c.JSON(http.StatusOK, updated)
}

func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	delete(books, id)
	c.Status(http.StatusNoContent)
}
