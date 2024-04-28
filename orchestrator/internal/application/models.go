package application

import (
	"Orchestrator/internal/http"
	"Orchestrator/internal/orchestrator"
	"database/sql"
)

type Application struct {
	DB           *sql.DB
	Orchestrator *orchestrator.Orchestrator
	HttpServer   *http.OrchServer
}
