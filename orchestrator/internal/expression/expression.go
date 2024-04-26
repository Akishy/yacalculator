package expression

import (
	"context"
	"errors"
	"fmt"
	"go/scanner"
	"go/token"
	"log"
	"orchestrator/internal/tokenizer"
)

func NewExpression(ctx context.Context, expression string) (*Expression, error) {
	expr := &Expression{
		//UserID:        0,
		//ExpressionID:  0,
		RawExpression: expression,
		//ASTExpression:
		FSet: token.NewFileSet(),
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
