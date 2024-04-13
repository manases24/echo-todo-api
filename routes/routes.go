package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mnsh5/todo-api/handlers"
	"github.com/mnsh5/todo-api/ping"
)

func SetupRoutes(e *echo.Echo) {
	task := handlers.TaskHandler{}

	// Grupo principal de la API
	api := e.Group("/api")

	// Subgrupo para la versión 1 de la API
	v1 := api.Group("/v1")

	v1.GET("", ping.Ping)

	// Task
	v1.GET("/tasks", task.FindAll)
	v1.GET("/task/:id", task.FindById)
	v1.POST("/task", task.Create)

}
