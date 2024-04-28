package sqlite

import (
	"context"
	"database/sql"
)

func createUsers(ctx context.Context, db *sql.DB) error {
	const usersTable = `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    password TEXT,
    is_admin BOOlEAN,
    amount_of_agents INTEGER CHECK(amount_of_agents >= 0),
    time_to_calc TEXT
	);`
	if _, err := db.ExecContext(ctx, usersTable); err != nil {
		return err
	}
	return nil
}

func createExpressions(ctx context.Context, db *sql.DB) error {
	const expressionsTable = `CREATE TABLE IF NOT EXISTS expressions (
        expression_id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        raw_expression TEXT,
        status INTEGER CHECK(status >= 0),
        result TEXT,
        FOREIGN KEY (user_id) REFERENCES users (id)
    );`
	if _, err := db.ExecContext(ctx, expressionsTable); err != nil {
		return err
	}

	return nil
}

func createSubExpressions(ctx context.Context, db *sql.DB) error {
	const subExpressionsTable = `CREATE TABLE IF NOT EXISTS sub_expressions (
    	subexpression_id INTEGER PRIMARY KEY,
    	expression_id INTEGER,
    	operator TEXT,
    	left_operand TEXT,
    	right_operand TEXT,
    	result TEXT,
    	FOREIGN KEY (subexpression_id) REFERENCES expressions (expression_id)
	);`
	if _, err := db.ExecContext(ctx, subExpressionsTable); err != nil {
		return err
	}

	return nil
}

func InitDB(ctx context.Context, db *sql.DB) error {
	err := createUsers(ctx, db)
	if err != nil {
		return err
	}

	err = createExpressions(ctx, db)
	if err != nil {
		return err
	}

	err = createSubExpressions(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
