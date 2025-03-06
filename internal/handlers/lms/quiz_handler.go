package lms

import (
	"echo-boilerplate/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetQuestionQuiz(c echo.Context) error {
	questionId := c.QueryParam("questiondetail_id")
	quizDetailId := c.QueryParam("quizdetail_id")

	if questionId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Question Id Tidak Boleh Kosong !",
		})
	}

	if quizDetailId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Quiz Detail Id Tidak Boleh Kosong !",
		})
	}

	questionIdInt, _ := strconv.Atoi(questionId)
	result, err := models.GetDetailQuestion(questionIdInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	quizDetailIdInt, _ := strconv.Atoi(quizDetailId)
	detailQuiz, err := models.GetDetailQuiz(quizDetailIdInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":         http.StatusOK,
		"message":        "Get Data Question Quiz",
		"detailQuestion": result,
		"detailQuiz":     detailQuiz,
	})
}
