package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	NoTelp       string `json:"no_telp"`
	Alamat       string `json:"alamat"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
}

type Claims struct {
	Username  string
	ExpiresAt time.Time
	jwt.RegisteredClaims
}
