package main

import (
	"github.com/ChaotenHG/filebased-template/logo"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	id := c.Param("id")

	value, err := logo.GetMerker(id)
	if err != nil {
		return err
	}

	return c.String(200, value)
}
