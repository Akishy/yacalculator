package subexpression

import (
	"go/ast"
	"go/constant"
	"go/token"
	"time"
)

type subExpressionStatus int

const (
	NEEDTOCALC subExpressionStatus = iota
	DONE
	STOPPED // todo: todo: добавить возможность останавливать вычисление подвыражения без блокировки агента
	FAILED
)

// SubExpression - минимальная единица, на которую можно поделить выражение. Отправлятся в таком виде Агенту калькулятора для вычисления.
type SubExpression struct {
	SubExpressionID     int                 // ID подвыражения, которое нужно для Агента
	ExpressionID        int                 // ID родительского выражения
	SubExpressionStatus subExpressionStatus // Статус подвыражения
	Left                ast.BasicLit        // Левый операнд
	Op                  token.Token         // Оператор
	Right               ast.BasicLit        // Правый оператор
	TimeToExec          time.Duration       // Время на выполнение подвыражения
	Result              constant.Value      // Результат вычисления

}
