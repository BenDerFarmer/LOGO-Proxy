package main

import (
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	name := c.Param("user")
	id := c.Param("id")
	return c.String(200, name+" "+id)
}
