package start

import (
	"echo-boilerplate/internal/auth"
	"echo-boilerplate/internal/customer"
	"echo-boilerplate/internal/product"
	"echo-boilerplate/pkg/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to Our API !")
	})
	v1 := e.Group("v1")
	ServiceAuth(v1)

	v1Protected := e.Group("v1", middleware.AuthMiddleware)
	ServiceProducts(v1Protected)
	ServiceCustomers(v1Protected)
}

func ServiceAuth(v1 *echo.Group) {
	v1.POST("/login", auth.LoginHandler)
	v1.POST("/register", auth.RegisterCustomer)

	v1.Group("", middleware.AuthMiddleware)
	{
		v1.GET("/profile", auth.ValidateToken)
	}
}

func ServiceProducts(v1 *echo.Group) {
	v1.GET("/products/get-products", product.GetProducts)
	v1.GET("/products/get-product-detail/:id", product.GetProductDetail)
	v1.POST("/products/add-product", product.AddProduct)
	v1.DELETE("/products/delete-product/:id", product.DeleteProduct)
}

func ServiceCustomers(v1 *echo.Group) {
	v1.GET("/customers/get-customers", customer.GetCustomers)
	v1.GET("/customers/get-customer-by-id/:id", customer.GetCustomerById)
	v1.DELETE("/customers/delete-customer/:id", customer.DeleteCustomerById)
}
