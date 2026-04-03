package models

import "time"

type Book struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Title      string  `gorm:"not null;size:255" json:"title" binding:"required,min=1"`
	AuthorID   uint    `gorm:"not null" json:"author_id" binding:"required,gt=0"`
	CategoryID uint    `gorm:"not null" json:"category_id" binding:"required,gt=0"`
	Price      float64 `gorm:"not null" json:"price" binding:"required,gt=0"`

	Author   Author   `gorm:"foreignKey:AuthorID" json:"author"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookInput struct {
	Title      string  `json:"title" binding:"required,min=1"`
	AuthorID   uint    `json:"author_id" binding:"required,gt=0"`
	CategoryID uint    `json:"category_id" binding:"required,gt=0"`
	Price      float64 `json:"price" binding:"required,gt=0"`
}
