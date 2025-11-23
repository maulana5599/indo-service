package customer

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func GetCustomers(c echo.Context) error {
	customerName := c.QueryParam("customer_name")

	result, err := FindAllCustomers(customerName)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, echo.Map{
			"status":  http.StatusOK,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Customers Successfully !",
		"data":    result,
	})
}

func GetCustomerById(c echo.Context) error {
	customerId, _ := strconv.Atoi(c.Param("id"))
	result, err := FindCustomerByID(customerId)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, echo.Map{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Customer By Id Success",
		"data":    result,
	})
}

func DeleteCustomerById(c echo.Context) error {
	customerId, _ := strconv.Atoi(c.Param("id"))

	validate := validation.Errors{
		"id": validation.Validate(customerId, validation.Required),
	}

	if err := validate.Filter(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err,
		})
	}

	err := DeleteCustomer(customerId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Delete Customer By Id Success",
	})
}
