package handlers

import (
	"echo-boilerplate/internal/entity"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	var loginRequest entity.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return err
	}

	// Set token expiration time
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := entity.Claims{
		Username:  "ahahaha",
		ExpiresAt: expirationTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(entity.JwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":       http.StatusOK,
		"message":      "Login successfully",
		"access_token": tokenString,
	})
}

func ValidateToken(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Token is valid",
		"token":   tokenString,
	})
}
