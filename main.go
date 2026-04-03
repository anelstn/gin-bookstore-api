package main

import (
	"git-practice-gin/config"
	"git-practice-gin/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBookByID)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)
	r.GET("/authors/:id", handlers.GetAuthorByID)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)
	r.GET("/categories/:id", handlers.GetCategoryByID)

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
