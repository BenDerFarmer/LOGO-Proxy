package main

import (
	"log"
	"os"
	"strings"

	"github.com/ChaotenHG/filebased-template/logo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	registerRoutes(e)

	var webui = os.Getenv("LOGO_PROXY_WEBUI")

	if strings.ToLower(webui) == "true" {
		e.Static("/", "./web/dist/")
		log.Println("Enabled Web UI")
	}

	if err := logo.LoadConfig(os.Getenv("LOGO_PROXY_URL"), os.Getenv("LOGO_PROXY_PASSWORD")); err != nil {
		log.Println(err)
	}

	var port = os.Getenv("LOGO_PROXY_PORT")

	if port == "" {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
