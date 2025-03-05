package lms

import (
	"echo-boilerplate/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SetConfigAnswerKey(c echo.Context) error {
	konfig := c.QueryParam("set-konfig")
	quizId := c.QueryParam("quiz_id")
	if konfig == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Set Konfig Tidak Boleh Kosong !",
		})
	}

	if quizId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Quiz ID Tidak Boleh Kosong !",
		})
	}

	konfigBool := false
	if konfig == "true" {
		konfigBool = true
	}

	quizInt, _ := strconv.Atoi(quizId)

	config.DB.Table("quiz_m").Where("quiz_id = ?", quizInt).Update("is_kuncijawaban", konfigBool)

	return c.JSON(http.StatusOK, echo.Map{
		"status":  200,
		"message": "Setting konfig berhasil dirubah !",
	})
}
