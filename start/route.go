package start

import (
	"echo-boilerplate/internal/handlers"
	"echo-boilerplate/pkg/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Selamat datang di service koperasi !")
	})
	v1 := e.Group("v1")
	ServiceAuth(v1)

	v1Protected := e.Group("v1", middleware.AuthMiddleware)
	ServiceUsers(v1Protected)
	ServiceJenisCicilan(v1Protected)
	ServiceCicilan(v1Protected)
	ServicePembayaranCicilan(v1Protected)
	ServiceMaster(v1Protected)
	ServiceLegal(v1Protected)
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
	v1.GET("/sysadmin/get-customer", handlers.GetCustomer)
	v1.GET("/sysadmin/search-customer", handlers.SearchCustomer)
}

func ServiceJenisCicilan(v1 *echo.Group) {
	v1.GET("/sysadmin/get-cicilan", handlers.GetJenisCicilan, middleware.AuthMiddleware)
	v1.GET("/sysadmin/get-detail-cicilan", handlers.GetJenisCicilanId, middleware.AuthMiddleware)
	v1.DELETE("/sysadmin/delete-jenis-cicilan", handlers.HapusCicilanById, middleware.AuthMiddleware)
	v1.POST("/sysadmin/add-jenis-cicilan", handlers.AddJenisCicilan, middleware.AuthMiddleware)
}

func ServiceCicilan(v1 *echo.Group) {
	v1.GET("/sysadmin/get-cicilan-siswa", handlers.GetDataCicilan)
	v1.POST("/sysadmin/get-cicilan-siswa-id", handlers.GetCicilanUser)
	v1.POST("/sysadmin/add-cicilan", handlers.AddCicilan)
	v1.DELETE("/sysadmin/batal-cicilan", handlers.BatalPengajuanCicilan)
}

func ServicePembayaranCicilan(v1 *echo.Group) {
	v1.GET("/sysadmin/get-pembayaran-cicilan-siswa/:pengajuan_id", handlers.GetPembayaranCicilanId)
	v1.GET("/sysadmin/get-pembayaran-detail-siswa/:pembayaran_id", handlers.GetPembayaranDetailId)
	v1.GET("/sysadmin/get-header-cicilan-siswa/:user_id", handlers.GetHeaderCicialnId)
	v1.GET("/sysadmin/generate-invoice", handlers.GenerateInvoice)
	v1.POST("/sysadmin/ubah-status-pembayaran-cicilan", handlers.UbahStatusPembayaranCicilan)
}

func ServiceMaster(v1 *echo.Group) {
	// Mata pelajaran.
	v1.GET("/sysadmin/get-mapel", handlers.GetMapel)
}

func ServiceLegal(v1 *echo.Group) {
	v1.GET("/sysadmin/get-grafik-jobs", handlers.GetGrafikJobs)
	v1.GET("/sysadmin/get-grafik-jobs-angkatan", handlers.GetGrafikJobsAngkatan)
}
