package services

import (
	"github.com/labstack/echo/v4"
	"github.com/mnsh5/todo-api/db"
)

type TaskService struct{}

func (t *TaskService) FindAll(c *echo.Context) {
	db := db.DB
}
