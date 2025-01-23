package handlers

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func GetDataCicilan(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Cicilan",
	})
}

func AddCicilan(c echo.Context) error {
	var request *entity.PengajuanCicilanRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	validationError := validation.ValidateStruct(request,
		validation.Field(&request.UserId, validation.Required),
		validation.Field(&request.NoKtp, validation.Required),
		validation.Field(&request.OrangTua, validation.Required),
	)

	if validationError != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": validationError.Error(),
		})
	}

	errCreate := models.AddCicilan(request)

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": errCreate.Error(),
		})
	}

	// Alur bisnisnya, ketika nanti melakukan pengajuan cicilan, maka akan langsung dibuatkan tagihannya
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Cicilan Berhasil Dibuat !",
		"payload": request,
	})
}

func BatalPengajuanCicilan(c echo.Context) error {
	pengajuanId, _ := strconv.Atoi(c.QueryParam("pengajuancicilan_id"))
	keterangan := c.QueryParam("keterangan")

	errBatal := models.BatalCicilan(pengajuanId, keterangan)

	if errBatal != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": errBatal.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Batal cicilan berhasil !",
	})
}

func GetCicilanUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("user_id"))
	result, err := models.GetCicilanUser(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var response []entity.PengajuanCicilanResponse
	for _, v := range result {
		response = append(response, entity.PengajuanCicilanResponse{
			PengajuancicilanId: v.PengajuancicilanId,
			UserId:             v.UserId,
			NoKtp:              v.NoKtp,
			Alamat:             v.Alamat,
			NoSiswa:            v.NoSiswa,
			Pekerjaan:          v.Pekerjaan,
			Orangtua:           v.Orangtua,
			NohpOrtu:           v.NohpOrtu,
			KontakDarurat:      v.KontakDarurat,
			Jaminan:            v.Jaminan,
			Keterangan:         nil,
			JeniscicilanId:     v.JeniscicilanId,
			CreatedAt:          v.CreatedAt.Format("2006-01-02"),
			UpdatedAt:          v.UpdatedAt.Format("2006-01-02"),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Cicilan User",
		"data":    response,
	})
}
