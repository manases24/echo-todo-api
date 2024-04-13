package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mnsh5/todo-api/db"
	"github.com/mnsh5/todo-api/routes"
)

func main() {
	e := echo.New()
	db.ConnectDB()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":2024"))
}
