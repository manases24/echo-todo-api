package routes

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {
	// Grupo principal de la API
	api := e.Group("/api")

	// Subgrupo para la versión 1 de la API
	v1 := api.Group("/v1")

}
