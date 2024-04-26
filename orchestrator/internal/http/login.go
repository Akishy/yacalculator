package http

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"orchestrator/internal/auth"
	"time"
)

func (osrv *OrchServer) login(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Throws unauthorized error
		user, err := auth.CheckLoginPassword(ctx, osrv.DB, username, password)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}
		if user == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
		}

		// Set custom claims
		claims := &jwtCustomClaims{
			username,
			user.IsAdmin,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(getJWTsecretFromEnv()))
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, echo.Map{"token": t})
	}

}
