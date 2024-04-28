package application

import (
	"Calculator/internal/calculator"
	"database/sql"
)

type Application struct {
	Calculator *calculator.Calculator
	DB         *sql.DB
}
