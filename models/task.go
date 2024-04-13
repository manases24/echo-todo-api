package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Done        bool   `json:"done" validate:"required"`
}
