package start

import (
	"echo-boilerplate/internal/handlers"
	"echo-boilerplate/pkg/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := e.Group("v1")
	ServiceAuth(v1)
	ServiceJenisCicilan(v1)
}

func ServiceAuth(v1 *echo.Group) {
	v1.POST("/login", handlers.LoginHandler)

	v1.Group("", middleware.AuthMiddleware)
	{
		v1.GET("/profile", handlers.ValidateToken)
	}
}

func ServiceJenisCicilan(v1 *echo.Group) {
	v1.GET("/sysadmin/get-cicilan", handlers.GetJenisCicilan, middleware.AuthMiddleware)
	v1.POST("/sysadmin/add-cicilan", handlers.AddJenisCicilan, middleware.AuthMiddleware)
}
