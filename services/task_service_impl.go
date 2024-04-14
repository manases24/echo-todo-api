package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mnsh5/todo-api/common"
	"github.com/mnsh5/todo-api/models"
	"gorm.io/gorm"
)

type TaskServiceImpl struct {
	Db *gorm.DB
}

func NewTaskServiceImpl(Db *gorm.DB) TaskService {
	return &TaskServiceImpl{Db: Db}
}

func (t *TaskServiceImpl) FindAll(c echo.Context) error {
	var tasks []models.Task
	// Intentar recuperar todos los usuarios de la base de datos
	if err := t.Db.Find(&tasks).Error; err != nil {
		errorMessage := fmt.Sprintf("Failed to fetch users: %s", err.Error())
		return common.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}
	response := map[string]interface{}{
		"status":  "success",
		"message": "All users retrieved successfully",
		"data":    &tasks,
	}
	return c.JSON(http.StatusOK, response)

}

func (t *TaskServiceImpl) FindById(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	if err := t.Db.First(&task, id).Error; err != nil {
		errorMessage := fmt.Sprintf("No task found with ID: %s", id)
		return common.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Task with the provided ID found",
		"data":    &task,
	}

	return c.JSON(http.StatusOK, response)
}

func (t *TaskServiceImpl) Create(c echo.Context) error {
	task := new(models.Task)

	if err := c.Bind(task); err != nil {
		errorMessage := fmt.Sprintf("Failed to parse request body: %s", err.Error())
		return common.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}

	if err := t.Db.Create(&task).Error; err != nil {
		errorMessage := fmt.Sprintf("Failed to create task: %s", err.Error())
		return common.ErrorHandler(c, http.StatusBadGateway, errorMessage, err)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Task has been created",
		"data":    &task,
	}

	return c.JSON(http.StatusCreated, response)
}

func (t *TaskServiceImpl) Update(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	if err := t.Db.First(&task, id).Error; err != nil {
		errorMessage := fmt.Sprintf("No task found with ID: %s", id)
		return common.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}

	updateTask := new(models.Task)
	if err := c.Bind(updateTask); err != nil {
		errorMessage := fmt.Sprintf("Failed to parse request body: %s", err.Error())
		return common.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}

	task.Name = updateTask.Name
	task.Description = updateTask.Description
	task.Done = updateTask.Done
	t.Db.Save(&task)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Task has been updated",
		"data":    &task,
	}

	return c.JSON(http.StatusOK, response)

}

func (t *TaskServiceImpl) Delete(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	if err := t.Db.First(&task, id).Error; err != nil {
		errorMessage := fmt.Sprintf("No task found with ID: %s", id)
		return common.ErrorHandler(c, http.StatusBadRequest, errorMessage, err)
	}

	t.Db.Delete(&task)
	response := map[string]interface{}{
		"status":  "success",
		"message": "Task has been deleted",
		"id":      &task.ID,
	}

	return c.JSON(http.StatusOK, response)
}
