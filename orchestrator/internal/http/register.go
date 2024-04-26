package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"orchestrator/internal/auth"
)

func (osrv *OrchServer) register(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		err := validateForbiddenSymbols(username)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		// password can contain any symbols
		err = validateLength(password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		exists, err := auth.CheckIfUserExists(ctx, osrv.DB, username)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}
		if exists {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": fmt.Sprintf("user %s already exists", username)})
		}

		user := auth.NewUser(username, password)

		userId, err := auth.InsertUser(ctx, osrv.DB, user)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}
		log.Printf("[INFO] register: user %v successefuly created", userId)
		return c.JSON(http.StatusCreated, []string{"OK"})
	}
}
