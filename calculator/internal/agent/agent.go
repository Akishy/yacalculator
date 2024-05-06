package agent

import (
	"Calculator/internal/task"
	"context"
	"database/sql"
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
		TaskChan:   make(chan task.Task),
		ResultChan: make(chan *Result),
	}
}

// calculateTask вычисляет выражение и отдаёт результат вычисления
func (a *Agent) calculateTask(left constant.Value, op token.Token, right constant.Value) constant.Value {
	log.Println("calculating...")
	out := constant.BinaryOp(left, op, right)
	log.Println(out.Kind())
	return out
}

func (a *Agent) sendHeartBeat() {
	timer := time.NewTicker(5 * time.Second) // Создаём таймер, который срабатывает каждые 5 секунд
	defer timer.Stop()                       // Убедимся, что таймер будет остановлен при выходе из функции

	for {
		select {
		case <-timer.C:
			log.Println(a.Status)
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
		return -1, err
	}

	a.Id = int(id)

	return id, nil
}

func (a *Agent) Run() {
	go a.sendHeartBeat() // отправляем статус агента в отдельной горутине
	for {
		select {
		case _task := <-a.TaskChan: // Когда прилетает таска
			log.Println("agent ", a.Id, "task catch")
			a.Status = WORKING // Меняем статус агента на "работаю"
			left := _task.Left
			op := _task.Operand
			right := _task.Right
			time.Sleep(_task.TimeToSolve * time.Nanosecond) // И считаем выражение с заданным таймаутом.
			out := a.calculateTask(left, op, right)
			result := &Result{
				Value:     out,
				ValueType: 0,
			}
			log.Println("out: /// ", out)
			a.ResultChan <- result // Отдаём результат в соответствующий канал
			log.Println("agent send out")
			_task.TaskStatus = task.DONE
			a.Status = ALIVE
		}
	}
}
