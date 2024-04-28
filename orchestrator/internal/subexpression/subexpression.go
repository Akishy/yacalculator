package subexpression

import (
	"go/ast"
	"go/constant"
	"go/token"
	"time"
)

func NewSubExpression(expressionID int, timeToExec time.Duration, left ast.BasicLit, op token.Token, right ast.BasicLit) *SubExpression {
	return &SubExpression{
		//SubExpressionID:     0,
		ExpressionID:        expressionID,
		SubExpressionStatus: 0,
		Left:                left,
		Op:                  op,
		Right:               right,
		TimeToExec:          timeToExec,
		Result:              nil,
	}
}

func (sexp *SubExpression) SendToAgent() constant.Value {

}
