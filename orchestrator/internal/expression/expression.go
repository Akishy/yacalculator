package expression

import (
	"Orchestrator/internal/tokenizer"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go/scanner"
	"go/token"
	"log"
)

func NewExpression(ctx context.Context, uid int, expression string) (*Expression, error) {
	expr := &Expression{
		//ExpressionID:  0,
		UserID:        uid,
		RawExpression: expression,
		//ASTExpression:
		FSet:   token.NewFileSet(),
		Status: NEED_TO_CALC,
		Result: nil,
	}

	err := expr.validateExpression()
	if err != nil {
		log.Printf("[ERROR] expression validation error: %v", err)
		return nil, err
	}
	log.Printf("[INFO] expression %v validation success", expr.ExpressionID)

	return expr, nil

}

func (expr *Expression) validateExpression() error {
	// берем наше исходное выражение
	src := []byte(expr.RawExpression)

	// инициализируем сканер
	var s scanner.Scanner

	file := expr.FSet.AddFile("", expr.FSet.Base(), len(src))
	s.Init(file, src, nil, 2)

	for {
		_, tok, _ := s.Scan()
		if tok == token.EOF {
			break
		}
		if !expr.isTokenAllowed(tok) {
			return errors.New(fmt.Sprintf("Found forbidden token %s in expression %s\nAllowed tokens: %v", tok.String(), expr.RawExpression, tokenizer.AllowedTokens))
		}

	}
	return nil
}

func (expr *Expression) validateToken(tok token.Token) bool {
	for _, allowedToken := range tokenizer.AllowedTokens {
		if allowedToken == tok {
			return true
		}
	}
	return false
}

func (expr *Expression) isTokenAllowed(tok token.Token) bool {
	for _, allowedToken := range tokenizer.AllowedTokens {
		if allowedToken == tok {
			return true
		}
	}
	return false
}

// InsertExpression - adds expression info to DB. Returns expression ID
func (expr *Expression) InsertExpression(ctx context.Context, db *sql.DB) (int64, error) {
	q := `INSERT INTO expressions (user_id, raw_expression, status, result) VALUES ($1, $2, $3, $4);`

	result, err := db.ExecContext(ctx, q, expr.UserID, expr.RawExpression, expr.Status, "")
	if err != nil {
		log.Printf("[ERROR] insertExpression: Error inserting new expression: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("[ERROR] insertExpression: Error getting integer generated by the database: %v", err)
		return 0, err
	}

	return id, nil
}
