package application

import (
	"Calculator/internal/calculator"
	"database/sql"
	"google.golang.org/grpc"
)

type Application struct {
	Calculator *calculator.Calculator
	DB         *sql.DB
	gRPCServer *grpc.Server
}
