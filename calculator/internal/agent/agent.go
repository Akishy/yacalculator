package agent

import (
	"context"
	"database/sql"
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"log"
	"time"
)

func NewAgent(ownerID int) *Agent {
	return &Agent{
		OwnerID:    ownerID,
		Status:     0,
		StatusChan: make(chan Status),
		TaskChan:   make(chan Task),
	}
}

// calculateTask вычисляет выражение и отдаёт результат вычисления
func (a *Agent) calculateTask(left ast.BasicLit, op token.Token, right ast.BasicLit) constant.Value {
	fmt.Println("calculating...")
	return constant.BinaryOp(constant.MakeFromLiteral(left.Value, left.Kind, 0), op, constant.MakeFromLiteral(right.Value, right.Kind, 0))
}

func (a *Agent) sendHeartBeat() {
	timer := time.NewTicker(5 * time.Second) // Создаём таймер, который срабатывает каждые 5 секунд
	defer timer.Stop()                       // Убедимся, что таймер будет остановлен при выходе из функции

	for {
		select {
		case <-timer.C:
			fmt.Println(a.Status)
			a.StatusChan <- a.Status // Отправляем текущее состояние агента
		}
	}
}

// Insert возвращает id агента
func (a *Agent) Insert(ctx context.Context, db *sql.DB) (int64, error) {
	q := `INSERT INTO agents (owner_id, status) VALUES ($1, $2);`

	result, err := db.ExecContext(ctx, q, a.OwnerID, a.Status)
	if err != nil {
		log.Printf("[ERROR] AgentInsert: error inserting agent: %v", err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("[ERROR] AgentInsert: error getting last insert id: %v", err)
	}

	return id, nil
}

func (a *Agent) Run(ctx context.Context) error {
	go a.sendHeartBeat()
	return nil
}
