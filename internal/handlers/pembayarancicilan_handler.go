package handlers

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
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

func UbahStatusPembayaranCicilan(c echo.Context) error {
	var request *entity.StatusPembayaranRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	validate := validation.ValidateStruct(request,
		validation.Field(&request.StatusPembayaran, validation.Required),
		validation.Field(&request.PembayarancicilanId, validation.Required),
	)

	if validate != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": validate.Error(),
		})
	}

	errUpdate := models.UbahStatusPembayaranCicilan(request)

	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": errUpdate.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":       http.StatusOK,
		"message":      "Status pembayaran berhasil dirubah !",
		"pembayaranId": request,
	})
}
