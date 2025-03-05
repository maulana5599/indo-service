package lms

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTopikPembelajaran(c echo.Context) error {
	roomId := c.QueryParam("room_id")

	if roomId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Room Id Tidak Boleh Kosong !",
		})
	}

	roomIdInt, err := strconv.Atoi(roomId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, _ := models.GetLearningTopic(roomIdInt)

	roomTopic, err := GetTopicRoom(roomIdInt)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get Data Topik Pembelajaran",
		"data":    result,
		"class":   roomTopic,
	})
}

func GetTopicRoom(roomId int) (entity.RoomTopic, error) {
	result, err := models.GetRoomTopic(roomId)

	return result, err
}
