package common

import (
	"github.com/labstack/echo/v4"
)

func ErrorHandler(c echo.Context, statusCode int, message string, err error) error {
	return c.JSON(statusCode, map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}
