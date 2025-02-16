package handlers

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
)

func GetHeaderCicialnId(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("user_id"))
	result, err := models.GetHeaderCicilanId(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Header Pembayaran Cicilan",
		"data":    result,
	})
}

func GetPembayaranCicilanId(c echo.Context) error {
	pengajuanId, _ := strconv.Atoi(c.Param("pengajuan_id"))
	result, err := models.GetPembayaranCicilanId(pengajuanId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Pembayaran Cicilan",
		"data":    result,
	})
}

func GetPembayaranDetailId(c echo.Context) error {
	pembayarancicilanId, _ := strconv.Atoi(c.Param("pembayaran_id"))
	result, err := models.GetPembayaranDetailId(pembayarancicilanId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Detail Pembayaran Cicilan",
		"data":    result,
	})
}

func UbahStatusPembayaranCicilan(c echo.Context) error {
	var request *entity.StatusPembayaranRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	validate := validation.ValidateStruct(request,
		validation.Field(&request.StatusPembayaran, validation.Required),
		validation.Field(&request.PembayarancicilanId, validation.Required),
	)

	if validate != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": validate.Error(),
		})
	}

	errUpdate := models.UbahStatusPembayaranCicilan(request)

	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": errUpdate.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":       http.StatusOK,
		"message":      "Status pembayaran berhasil dirubah !",
		"pembayaranId": request,
	})
}

func GenerateInvoice(c echo.Context) error {

	// TODO: generate invoice
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "INVOICE")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 220, 255)
	pdf.Cell(90, 7, "Description")
	pdf.Cell(30, 7, "Quantity")
	pdf.Cell(30, 7, "Unit Price")
	pdf.Cell(40, 7, "Total")
	pdf.Ln(7)

	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Generate Invoice",
	})
}
