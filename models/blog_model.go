package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title       string `json:"title"`
	Datetime    string `json:"datetime"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
