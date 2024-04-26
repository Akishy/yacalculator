package http

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

func register(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		return nil
	}
}
