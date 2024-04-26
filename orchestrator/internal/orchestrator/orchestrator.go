package orchestrator

import (
	"context"
	"database/sql"
)

func RunOrchestrator(ctx context.Context, database *sql.DB) *Orchestrator {
	orchestrator := new(Orchestrator)
	err := database.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	orchestrator.db = database
	return orchestrator
}

// TODO: принять выражение от web'а, спарсить его на токены и валидировать, записать выражение в базу данных если оно нормальное. Произвести разбиение на подвыражения и отправить сервису калькулятор подвыражения на вычисление.

// TODO: Разбить выражение на подвыражения.
