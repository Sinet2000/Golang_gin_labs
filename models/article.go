package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
