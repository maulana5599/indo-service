package handlers

import (
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetGrafikJobs(c echo.Context) error {
	result, _ := models.GetGrafikJobs()

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Grafik Jobs",
		"data":    result,
	})
}

func GetGrafikJobsAngkatan(c echo.Context) error {
	perusahaanId := c.QueryParam("perusahaan_id")

	if perusahaanId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Perusahaan Id Tidak Boleh Kosong !",
		})
	}

	perusahaanIdInt, _ := strconv.Atoi(perusahaanId)

	result, _ := models.GetGrafikJobsAngkatan(perusahaanIdInt)

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Grafik Jobs",
		"data":    result,
	})
}
