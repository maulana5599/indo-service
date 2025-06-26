package mensetsu

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMensetsu(c echo.Context) error {
	page := c.QueryParam("page")
	perPage := c.QueryParam("per_page")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Mensetsu",
		"page":     page,
		"per_page": perPage,
	})
}

func GetDetailMensetsuNyId(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Detail Mensetsu",
	})
}
