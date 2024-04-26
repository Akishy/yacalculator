package expression

import (
	"go/ast"
	"go/token"
)

// Expression - основная структура выражения, которое отправил на вычисление пользователь
type Expression struct {
	ExpressionID  int            // ID выражения
	UserID        int            // ID пользователя
	RawExpression string         // Выражение, отправленное пользователем
	ASTExpression ast.Expr       // Бинарное дерево подвыражений, сформированное из исходного выражения
	FSet          *token.FileSet // token Fileset
}
