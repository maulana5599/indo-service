package product

import (
	"math/rand"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	searchParam := c.QueryParam("search")

	result := FindAllProducts(searchParam)

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Products Success",
		"data":    result,
	})
}

func GetProductDetail(c echo.Context) error {
	productId := c.Param("id")

	if productId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Product ID is required",
		})
	}

	result, err := FindProductByID(productId)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, echo.Map{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Product Detail Success",
		"data":    result,
	})
}

func AddProduct(c echo.Context) error {
	request := new(ProductRequest)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	validate := validation.Errors{
		"product_name":      validation.Validate(request.ProductName, validation.Required),
		"product_price":     validation.Validate(request.ProductPrice, validation.Required),
		"brand":             validation.Validate(request.Brand, validation.Required),
		"product_image_url": validation.Validate(request.ProductImageUrl, validation.Required),
	}

	if err := validate.Filter(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err,
		})
	}

	productId := rundDigits(10)

	payload := Product{
		ID:              productId,
		ProductName:     request.ProductName,
		ProductPrice:    request.ProductPrice,
		Brand:           request.Brand,
		ProductImageUrl: request.ProductImageUrl,
		ProductInfo:     request.ProductInfo,
	}

	errCreate := CreateProduct(payload)

	if errCreate != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": errCreate.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Add Product Success",
		"data":    payload,
	})
}

func rundDigits(n int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}

func UpdateProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Update Product Success",
	})
}

func DeleteProduct(c echo.Context) error {
	productId := c.Param("id")

	validation := validation.Errors{
		"id": validation.Validate(productId, validation.Required),
	}

	if err := validation.Filter(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err,
		})
	}

	errDelete := DeleteProductById(productId)

	if errDelete != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": errDelete.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Delete Product Success",
	})
}
