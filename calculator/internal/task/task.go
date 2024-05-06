package task

import (
	"context"
	"database/sql"
	"go/constant"
	"go/token"
	"log"
	"time"
)

func NewTask(left, right constant.Value, operand token.Token, timeToSolve time.Duration) *Task {
	return &Task{
		Left:        left,
		Right:       right,
		Operand:     operand,
		TimeToSolve: timeToSolve,
		TaskStatus:  NEED_TO_CALC,
	}
}

func (t *Task) Insert(ctx context.Context, db *sql.DB) (int64, error) {
	q := `INSERT INTO task (left, right, operand, time_to_solve) VALUES ($1, $2, $3, $4)`

	result, err := db.ExecContext(ctx, q, t.Left, t.Right, t.Operand, t.TimeToSolve)
	if err != nil {
		log.Printf("[ERROR] TaskInsert: error inserting task into DB: %v", err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("[ERROR] TaskInsert: error getting last insert ID: %v", err)
		return -1, err
	}
	t.Id = int(id)
	return id, err
}
