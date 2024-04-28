package expression

import (
	"Orchestrator/internal/tokenizer"
	"context"
	"go/token"
	"testing"
)

// TestNewExpression проверяет создание нового объекта Expression и его валидацию.
func TestNewExpression(t *testing.T) {
	ctx := context.Background()
	uid := 1
	expression := "2 + 2"

	// Имитация разрешенных токенов для теста
	tokenizer.AllowedTokens = []token.Token{token.INT, token.ADD}

	expr, err := NewExpression(ctx, uid, expression)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if expr.UserID != uid {
		t.Errorf("Expected UserID %d, got %d", uid, expr.UserID)
	}

	if expr.RawExpression != expression {
		t.Errorf("Expected RawExpression %s, got %s", expression, expr.RawExpression)
	}

	if expr.Status != NEED_TO_CALC {
		t.Errorf("Expected Status %v, got %v", NEED_TO_CALC, expr.Status)
	}

	// Проверка на невалидное выражение
	invalidExpression := "2 + unknown"
	_, err = NewExpression(ctx, uid, invalidExpression)
	if err == nil {
		t.Error("Expected error for invalid expression, got nil")
	}
}

// TestIsTokenAllowed проверяет разрешен ли токен.
func TestIsTokenAllowed(t *testing.T) {
	// Имитация разрешенных токенов
	tokenizer.AllowedTokens = []token.Token{token.INT, token.ADD}

	expr := &Expression{
		RawExpression: "2 + 2",
	}

	// Проверяем разрешенные токены
	if !expr.isTokenAllowed(token.INT) {
		t.Error("Expected token INT to be allowed")
	}

	if !expr.isTokenAllowed(token.ADD) {
		t.Error("Expected token ADD to be allowed")
	}

	// Проверяем неразрешенный токен
	if expr.isTokenAllowed(token.SUB) {
		t.Error("Expected token SUB to be disallowed")
	}
}
