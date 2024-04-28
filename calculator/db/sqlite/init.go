package sqlite

import (
	"context"
	"database/sql"
)

func createAgents(ctx context.Context, db *sql.DB) error {
	const agentsTable = `CREATE TABLE IF NOT EXISTS agents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    owner_id INTEGER CHECK (owner_id >= 0),
    status INTEGER
);`
	if _, err := db.ExecContext(ctx, agentsTable); err != nil {
		return err
	}

	return nil
}

func createTasks(ctx context.Context, db *sql.DB) error {
	const tasksTable = `CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    status INTEGER,
    time_to_calc TEXT,
    result TEXT
);`
	if _, err := db.ExecContext(ctx, tasksTable); err != nil {
		return err
	}

	return nil
}

func InitDB(ctx context.Context, db *sql.DB) error {
	err := createAgents(ctx, db)
	if err != nil {
		return err
	}
	err = createTasks(ctx, db)
	if err != nil {
		return err
	}
	return nil
}
