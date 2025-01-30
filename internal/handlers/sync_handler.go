package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func SyncHandlerUser(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
	})
}