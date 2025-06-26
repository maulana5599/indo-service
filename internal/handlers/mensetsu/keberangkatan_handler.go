package mensetsu

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetKeberangkatanSiswa(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Keberangkatan Siswa",
	})
}

func GetKeberangkatanDokumen(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Keberangkatan Dokumen",
	})
}
