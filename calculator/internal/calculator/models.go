package calculator

import (
	"Calculator/internal/agent"
	"database/sql"
)

type Calculator struct {
	DB     *sql.DB
	Agents map[int]*agent.Agent
}
