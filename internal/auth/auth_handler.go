package auth

import (
	"net/http"
	"time"

	customerEntity "echo-boilerplate/internal/customer"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c echo.Context) error {
	var loginRequest LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return err
	}

	result, errPassword := GetUsername(loginRequest.Email)

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
	expirationTime := time.Now().Add(3600 * time.Minute)
	claims := Claims{
		Username:  loginRequest.Email,
		ExpiresAt: expirationTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Set Cookie
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	// Get Role User
	role, _ := GetRole(result.Id)
	return c.JSON(http.StatusOK, echo.Map{
		"status":       http.StatusOK,
		"message":      "Login successfully !",
		"access_token": tokenString,
		"profile": echo.Map{
			"name":  result.Name,
			"email": result.Email,
		},
		"role": role,
	})
}

func RegisterCustomer(c echo.Context) error {
	request := RegisterRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validate := validation.Errors{
		"email":         validation.Validate(request.Email, validation.Required),
		"password":      validation.Validate(request.Password, validation.Required),
		"tanggal_lahir": validation.Validate(request.TanggalLahir, validation.Required),
		"tempat_lahir":  validation.Validate(request.TempatLahir, validation.Required),
	}

	if err := validate.Filter(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err,
		})
	}

	// Check Unique Email
	if resultEmail, _ := CheckEmail(request.Email); resultEmail.Email != "" {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Email already exists",
		})
	}

	birthdateParse, err := time.Parse("2006-01-02", request.TanggalLahir)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	passwordHash := []byte(request.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordHash, bcrypt.DefaultCost)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	payloadUser := Users{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	payloadCustomer := customerEntity.Customer{
		Alamat:       request.Alamat,
		NoTelp:       request.NoTelp,
		TempatLahir:  request.TempatLahir,
		TanggalLahir: birthdateParse,
	}

	errCreate := CreateNewUser(payloadUser, payloadCustomer)

	if errCreate != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": errCreate.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Register Customer Success",
		"data":    payloadCustomer,
		"user":    payloadUser,
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
