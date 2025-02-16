package handlers

import (
	"echo-boilerplate/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMapel(c echo.Context) error {
	result, _ := models.GetMapel()

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Mapel",
		"data":    result,
	})
}
