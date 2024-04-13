package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `gorm:"not null" validate:"required,min=1,max=200" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Done        bool   `gorm:"not null" json:"done"`
}
