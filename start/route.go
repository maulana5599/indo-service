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

	v1 := e.Group("v1", middleware.AuthMiddleware)
	ServiceAuth(v1)
	ServiceUsers(v1)
	ServiceJenisCicilan(v1)
	ServiceCicilan(v1)
	ServicePembayaranCicilan(v1)
}

func ServiceAuth(v1 *echo.Group) {
	v1.POST("/login", handlers.LoginHandler)

	v1.Group("", middleware.AuthMiddleware)
	{
		v1.GET("/profile", handlers.ValidateToken)
	}
}

func ServiceUsers(v1 *echo.Group) {
	v1.GET("/sysadmin/get-users", handlers.GetSiswa)
	v1.GET("/sysadmin/get-users/:id", handlers.GetSiswaById)
	v1.GET("/sysadmin/sync-users", handlers.SyncUsers)
}

func ServiceJenisCicilan(v1 *echo.Group) {
	v1.GET("/sysadmin/get-cicilan", handlers.GetJenisCicilan, middleware.AuthMiddleware)
	v1.POST("/sysadmin/add-cicilan", handlers.AddJenisCicilan, middleware.AuthMiddleware)
}

func ServiceCicilan(v1 *echo.Group) {
	v1.GET("/sysadmin/get-cicilan-siswa", handlers.GetDataCicilan)
	v1.POST("/sysadmin/get-cicilan-siswa-id", handlers.GetCicilanUser)
	v1.POST("/sysadmin/add-cicilan", handlers.AddCicilan)
	v1.DELETE("/sysadmin/batal-cicilan", handlers.BatalPengajuanCicilan)
}

func ServicePembayaranCicilan(v1 *echo.Group) {
	v1.GET("/sysadmin/get-pembayaran-cicilan-siswa/:user_id", handlers.GetPembayaranCicilanId)
	v1.POST("/sysadmin/ubah-status-pembayaran-cicilan", handlers.UbahStatusPembayaranCicilan)
}
