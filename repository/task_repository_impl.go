package repository

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mnsh5/todo-api/models"
	"github.com/mnsh5/todo-api/utils"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(DB *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: DB}
}

func (t *TaskRepositoryImpl) FindAll(c echo.Context) error {
	var tasks []models.Task
	// Intentar recuperar todos los usuarios de la base de datos
	if err := t.Db.Find(&tasks).Error; err != nil {
		errorMessage := fmt.Sprintf("Failed to fetch users: %s", err.Error())
		return utils.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}
	response := map[string]interface{}{
		"status":  "success",
		"message": "All users retrieved successfully",
		"data":    &tasks,
	}
	return c.JSON(http.StatusOK, response)

}

func (t *TaskRepositoryImpl) FindById(c echo.Context) error {
	id := c.Param("id")
	var task []models.Task

	if err := t.Db.Find(&task, id).Error; err != nil {
		errorMessage := fmt.Sprintf("No user found with ID: %s", id)
		return utils.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "User with the provided ID found",
		"data":    &task,
	}

	return c.JSON(http.StatusOK, response)

}

func (t *TaskRepositoryImpl) Save(c echo.Context) {
	// Implementación de la función Save
}

func (t *TaskRepositoryImpl) Update(c echo.Context) {
	// Implementación de la función Update
}

func (t *TaskRepositoryImpl) Delete(c echo.Context) {
	// Implementación de la función Delete
}
