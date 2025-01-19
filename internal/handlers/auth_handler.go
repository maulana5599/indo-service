package handlers

import (
	"echo-boilerplate/internal/entity"
	"echo-boilerplate/internal/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c echo.Context) error {
	var loginRequest entity.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return err
	}

	result, errPassword := models.GetUsername(loginRequest.Email)

	if errPassword != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized !",
		})
	}

	hashedPassword := []byte(result.Password)
	passwordRequest := []byte(loginRequest.Password)
	errCheck := bcrypt.CompareHashAndPassword(hashedPassword, passwordRequest)

	if errCheck != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}

	// Set token expiration time
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := entity.Claims{
		Username:  loginRequest.Email,
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

	// Set Cookie
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

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
