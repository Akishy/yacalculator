package agent

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"time"
)

func NewAgent(ownerID int) *Agent {
	return &Agent{
		OwnerID:    ownerID,
		Status:     0,
		StatusChan: make(chan Status),
	}
}

// calculateTask вычисляет выражение и отдаёт
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
