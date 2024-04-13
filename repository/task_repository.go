package repository

import (
	"github.com/labstack/echo/v4"
)

type TaskRepository interface {
	Save(c echo.Context)
	Update(c echo.Context)
	Delete(c echo.Context)
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
}
