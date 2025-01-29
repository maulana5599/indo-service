package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("secret")

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenString := ctx.Request().Header.Get("Authorization")
		if tokenString == "" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized !",
			})
		}

		stringCut := strings.Replace(tokenString, "Bearer ", "", -1)

		err := VerifyToken(stringCut)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
		}

		return next(ctx)
	}
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
