package main

import (
	"echo-boilerplate/config"
	"echo-boilerplate/start"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// Load Configuration with ENV.
	loadEnv()
	portApp := os.Getenv("APP_PORT")

	// Load Database.
	config.DatabaseConnection()

	// Load Route.
	start.Route(e)
	e.Logger.Fatal(e.Start(":" + portApp))
}

func loadEnv() {
	env := godotenv.Load(".env")
	if env != nil {
		log.Println("Error loading .env file")
		panic("Error loading .env file")
	}
}
