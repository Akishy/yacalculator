package agent

import (
	"go/ast"
	"go/token"
	"time"
)

type Status int

const (
	ALIVE   Status = iota // Значит что можно поручить задачку
	WORKING               // Значит что живой, но работает
)

type TaskStatus int

const (
	NEED_TO_CALC TaskStatus = iota
	CALCULATING
	DONE
	STOPPED
)

type Agent struct {
	Id         int         // Идентификатор агента
	OwnerID    int         // Идентификатор владельца агента
	Status     Status      // Статус агента
	StatusChan chan Status // Канал для отправки статуса калькулятора
	TaskChan   chan Task
}

type SubExpression struct {
	Left       ast.BasicLit
	Op         token.Token
	Right      ast.BasicLit
	TimeToExec time.Duration
}

// Task - задание, которое отправляется вычислителю (агенту) на выполнение. Содержит в себе подвыражение, а так же дополнительные поля, такие как идентификатор задания, время выполнения выражения, и статус задания.
type Task struct {
	TaskID  int
	SubExpr SubExpression
	Status  TaskStatus
}
