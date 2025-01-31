package handlers

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/labstack/echo/v4"
)

func GetJenisCicilan(c echo.Context) error {
	result, err := models.GetJenisCicilan()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var response []entity.JenisCicilanResponse
	for _, v := range result {
		response = append(response, entity.JenisCicilanResponse{
			JenispinjamanId: v.JenispinjamanId,
			NamaCicilan:     v.NamaCicilan,
			PokokCicilan:    v.PokokCicilan,
			TotalAngsuran:   v.TotalAngsuran,
			JumlahAngsuran:  v.JumlahAngsuran,
			MarginCicilan:   v.MarginCicilan,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Jenis Cicilan",
		"data":    response,
	})
}

func GetJenisCicilanId(c echo.Context) error {
	cicilanId := c.QueryParam("cicilan_id")
	cicilanIdInt, _ := strconv.Atoi(cicilanId)

	validate := validation.Validate(cicilanIdInt, validation.Required)

	if validate != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Jenis Cicilan Id Tidak Boleh Kosong !",
		})
	}

	result, err := models.GetJenisCicilanId(cicilanIdInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	response := entity.JenisCicilanResponse{
		JenispinjamanId: result.JenispinjamanId,
		NamaCicilan:     result.NamaCicilan,
		PokokCicilan:    result.PokokCicilan,
		TotalAngsuran:   result.TotalAngsuran,
		JumlahAngsuran:  result.JumlahAngsuran,
		MarginCicilan:   result.MarginCicilan,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Jenis Cicilan By Id",
		"data":    response,
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

func HapusCicilanById(c echo.Context) error {
	cicilanId := c.QueryParam("cicilan_id")
	cicilanIdInt, _ := strconv.Atoi(cicilanId)

	err := models.DeleteCicilan(cicilanIdInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Hapus Jenis Cicilan Berhasil !",
	})
}
