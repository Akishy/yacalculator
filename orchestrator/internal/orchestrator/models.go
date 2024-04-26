package orchestrator

import "database/sql"

type Orchestrator struct {
	db *sql.DB
}
