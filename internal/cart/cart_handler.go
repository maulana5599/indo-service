package cart

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCart(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Cart Success",
	})
}

func GetCartByCustomer(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Cart By Customer Success",
	})
}
