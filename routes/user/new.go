package main

import (
	"github.com/labstack/echo/v4"
)

func Post_new(c echo.Context) error {
	return c.String(200, "post new user")
}
