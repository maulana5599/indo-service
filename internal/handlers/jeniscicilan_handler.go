package handlers

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/labstack/echo/v4"
)

func GetJenisCicilan(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Jenis Cicilan",
	})
}

func AddJenisCicilan(c echo.Context) error {
	request := new(entity.JenisCicilanRequest)
	if err := c.Bind(&request); err != nil {
		return err
	}

	validName := validation.Validate(request.NamaCicilan, validation.Required, validation.Length(1, 50), models.UniqueNameDB())

	if validName != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": validName.Error(),
		})
	}

	validationError := validation.ValidateStruct(request,
		validation.Field(&request.PokokCicilan, validation.Required),
		validation.Field(&request.JumlahAngsuran, validation.Required),
	)

	if validationError != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": validationError.Error(),
		})
	}

	payload := &entity.JenisCicilan{
		NamaCicilan:    request.NamaCicilan,
		PokokCicilan:   request.PokokCicilan,
		TotalAngsuran:  request.TotalAngsuran,
		JumlahAngsuran: request.JumlahAngsuran,
		MarginCicilan:  request.MarginCicilan,
	}

	models.AddJenisCicilan(payload)

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Tambah Jenis Cicilan Berhasil !",
	})
}
