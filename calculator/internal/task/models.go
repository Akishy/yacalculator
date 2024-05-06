package task

import (
	"go/constant"
	"go/token"
	"time"
)

type taskStatus int

const (
	NEED_TO_CALC taskStatus = iota
	CALCULATING
	DONE
	STOPPED
)

// Task - задание, которое отправляется вычислителю (агенту) на выполнение. Содержит в себе подвыражение, а так же дополнительные поля, такие как идентификатор задания, время выполнения выражения, и статус задания.
type Task struct {
	Id          int
	Left        constant.Value
	Right       constant.Value
	Operand     token.Token
	TimeToSolve time.Duration
	result      constant.Value
	TaskStatus  taskStatus
}
