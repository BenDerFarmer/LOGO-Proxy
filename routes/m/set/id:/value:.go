package routes

import (
	"github.com/ChaotenHG/filebased-template/logo"
	"github.com/labstack/echo/v4"
)

func Post_set(c echo.Context) error {
	id := c.Param("id")
	value := c.Param("value")

	if err := logo.SetMerker(id, value); err != nil {
		return err
	}

	return c.String(200, id+value)
}
