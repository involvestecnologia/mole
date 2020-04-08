package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewHealthHandler(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
