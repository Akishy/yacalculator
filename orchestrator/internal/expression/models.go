package expression

import (
	"go/ast"
	"go/token"
	"reflect"
)

type ExprResult interface {
	String() string
	Type() reflect.Type
}

type expressionStatus int

const (
	NEED_TO_CALC expressionStatus = iota
	CALCULATING
	DONE
	STOPPED // todo: добавить возможность останавливать вычисление
	FAILED
)

// Expression - основная структура выражения, которое отправил на вычисление пользователь
type Expression struct {
	ExpressionID  int              `json:"expression_ID,omitempty"`  // ID выражения
	UserID        int              `json:"user_ID,omitempty"`        // ID пользователя
	RawExpression string           `json:"raw_expression"`           // Выражение, отправленное пользователем
	ASTExpression ast.Expr         `json:"ast_expression,omitempty"` // Бинарное дерево подвыражений, сформированное из исходного выражения
	FSet          *token.FileSet   `json:"f_set,omitempty"`          // token Fileset
	Status        expressionStatus `json:"status,omitempty"`
	Result        ExprResult       `json:"result,omitempty"`
}
