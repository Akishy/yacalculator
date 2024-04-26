package tokenizer

import "go/token"

var AllowedTokens = []token.Token{
	token.ADD,
	token.SUB,
	token.MUL,
	token.QUO,
	token.INT,
	token.FLOAT,
}
