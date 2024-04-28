package subexpression

import (
	"go/ast"
	"go/token"
	"testing"
	"time"
)

// TestNewSubExpression проверяет создание нового SubExpression.
func TestNewSubExpression(t *testing.T) {
	expressionID := 1
	timeToExec := 5 * time.Second

	left := ast.BasicLit{
		Kind:  token.INT,
		Value: "2",
	}

	right := ast.BasicLit{
		Kind:  token.INT,
		Value: "3",
	}

	op := token.ADD

	subExp := NewSubExpression(expressionID, timeToExec, left, op, right)

	// Проверяем поля SubExpression
	if subExp.ExpressionID != expressionID {
		t.Errorf("expected ExpressionID %d, got %d", expressionID, subExp.ExpressionID)
	}

	if subExp.TimeToExec != timeToExec {
		t.Errorf("expected TimeToExec %v, got %v", timeToExec, subExp.TimeToExec)
	}

	if subExp.Left.Kind != token.INT || subExp.Left.Value != "2" {
		t.Errorf("expected Left BasicLit token.INT with value '2', got %v with value '%s'", subExp.Left.Kind, subExp.Left.Value)
	}

	if subExp.Right.Kind != token.INT || subExp.Right.Value != "3" {
		t.Errorf("expected Right BasicLit token.INT with value '3', got %v with value '%s'", subExp.Right.Kind, subExp.Right.Value)
	}

	if subExp.Op != op {
		t.Errorf("expected Operator %v, got %v", op, subExp.Op)
	}
}
