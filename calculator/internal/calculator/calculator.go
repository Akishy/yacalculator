package calculator

import (
	"context"
	"database/sql"
)

func New(db *sql.DB) *Calculator {
	return &Calculator{
		DB: db,
	}
}

func (calc *Calculator) Start(ctx context.Context) {
	err := calc.DB.PingContext(ctx)
	if err != nil {
		panic(err)
	}
}
