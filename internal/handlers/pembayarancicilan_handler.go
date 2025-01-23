package handlers

import (
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPembayaranCicilanId(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("user_id"))
	result, err := models.GetPembayaranCicilanId(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Pembayaran Cicilan",
		"data":    result,
	})
}
