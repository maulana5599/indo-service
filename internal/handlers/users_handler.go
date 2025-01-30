package handlers

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetSiswa(c echo.Context) error {
	result, _ := models.GetUsers()
	var responseUser []entity.ResponseUsers
	for _, v := range result {
		responseUser = append(responseUser, entity.ResponseUsers{
			Id:        v.Id,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Siswa",
		"data":    responseUser,
	})
}

func GetSiswaById(c echo.Context) error {
	usersId, _ := strconv.Atoi(c.Param("id"))

	result, _ := models.GetUsersById(usersId)

	if result.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{
			"status":  http.StatusNotFound,
			"message": "Data Tidak Ditemukan !",
		})
	}

	responseUser := entity.ResponseUsers{
		Id:        result.Id,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Siswa By Id",
		"data":    responseUser,
	})
}

func SyncUsers(c echo.Context) error {

	var url string = "https://jsonplaceholder.typicode.com/posts"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	defer resp.Body.Close()

	// Read the response body
	resultData, err := io.ReadAll(resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Ubah ke bentuk json
	var jsonData []map[string]interface{}

	if err := json.Unmarshal(resultData, &jsonData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Timpa data
	var users []entity.Users
	for _, v := range jsonData {
		users = append(users, entity.Users{
			Name:  v["title"].(string),
			Email: v["title"].(string),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Sync Users",
		"data":    users,
	})
}

func GetCustomer(c echo.Context) error {
	result, _ := models.GetCustomerAll()

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get All Users",
		"data":    result,
	})
}

func SearchCustomer(c echo.Context) error {
	name := c.QueryParam("nama_customer")
	result, _ := models.SearchCustomer(name)

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get All Users",
		"data":    result,
	})
}
