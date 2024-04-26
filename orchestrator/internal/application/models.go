package application

import (
	"database/sql"
	"orchestrator/internal/http"
	"orchestrator/internal/orchestrator"
)

type Application struct {
	DB           *sql.DB
	Orchestrator *orchestrator.Orchestrator
	HttpServer   *http.OrchServer
}
