package auth

import (
	"time"
)

type Users struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	RoleId int
	UserId int
}

func (Users) TableName() string {
	return "users"
}

var JwtKey = []byte("secret")

type EmailDatabase struct {
	Id       int
	Name     string
	Email    string
	Password string
}
