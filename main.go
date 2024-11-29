package main

import (
	"log"
	"os"

	"github.com/ChaotenHG/filebased-template/logo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	registerRoutes(e)

	if err := logo.LoadConfig(os.Getenv("LOGO_PROXY_URL"), os.Getenv("LOGO_PROXY_PASSWORD")); err != nil {
		log.Println(err)
	}

	var port = os.Getenv("LOGO_PROXY_PORT")

	if port == "" {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
