package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secret")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailDatabase struct {
	Email    string
	Password string
}

type Claims struct {
	Username  string
	ExpiresAt time.Time
	jwt.RegisteredClaims
}
