package http

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"math/rand"
	"os"
	"strings"
)

type OrchServer struct {
	Echo *echo.Echo
	DB   *sql.DB
}

func NewOrchServer(db *sql.DB) *OrchServer {
	return &OrchServer{
		Echo: echo.New(),
		DB:   db,
	}
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func (osrv *OrchServer) RunHttpServer(ctx context.Context) {
	e := osrv.Echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/login", osrv.login(ctx))
	e.POST("/register", osrv.register(ctx))

	api := e.Group("/api")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(getJWTsecretFromEnv()),
	}

	api.Use(echojwt.WithConfig(config))

	api.GET("/expressions", getAllExpressions)
	api.POST("/expressions", createExpression)
	api.DELETE("/expressions/:id", deleteExpressionById)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}

// validate Functions

func validateLength(value string) error {
	if len(value) == 0 {
		log.Println("[DEBUG] validateLength: length is zero")
		return errors.New("validate error")
	}
	return nil
}
func validateForbiddenSymbols(values ...string) error {
	const forbidden = "_~:/?#[]@!$&'()*+,;="
	for _, value := range values {

		err := validateLength(value)
		if err != nil {
			return err
		}

		for _, symbol := range value {
			if strings.ContainsRune(forbidden, symbol) {
				log.Println("[DEBUG] validateForbiddenSymbols: forbidden", symbol)
				return errors.New("validate error")
			}
		}
	}
	return nil
}

// secret functions

func getJWTsecretFromEnv() string {
	envSecret := os.Getenv("JWT_SECRET")

	if envSecret == "" {
		var randomString string
		for i := 0; i < 32; i++ {
			randomString += string(byte(rand.Intn(26) + 'a'))
		}

		err := os.Setenv("JWT_SECRET", randomString)
		if err != nil {
			log.Fatal("[WARNING: set JWT_SECRET to env failed: ]", err)
		}

		return randomString

	}
	return envSecret
}
