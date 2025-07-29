package lexer

import (
	"encoding/json"
	"os"
	"regexp"
)

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	IDENT   TokenType = "IDENT"
	INT     TokenType = "INT"
	STRING  TokenType = "STRING"

	// Keywords
	PACKAGE TokenType = "PACKAGE"
	FUNC    TokenType = "FUNC"
	VAR     TokenType = "VAR"
	TYPE    TokenType = "TYPE"
	IF      TokenType = "IF"
	ELSE    TokenType = "ELSE"
	FOR     TokenType = "FOR"
	RETURN  TokenType = "RETURN"
	// Add more...

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	ASTER  = "*"
	SLASH  = "/"
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NEQ    = "!="
	// Add more...

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

var TokenMap map[TokenType]TokenType = map[TokenType]TokenType{
	"package": PACKAGE,
	"func":    FUNC,
	"var":     VAR,
}

type Token struct {
	Type    TokenType
	Literal string
	Pos     int // optional: byte offset
}

func NewToken(t TokenType, lit string, pos int) Token {
	return Token{
		Type:    t,
		Literal: lit,
		Pos:     pos,
	}
}

type TokenResult struct {
	Token  Token
	IsZero bool
	Error  error
}

func ResultToken(t Token, err error) TokenResult {
	return TokenResult{
		Token: t,
		Error: err,
	}
}

func ResultTokenEmpty() TokenResult {
	return TokenResult{
		IsZero: true,
	}
}

type TokenDef struct {
	Id            string `json:"id"`
	Pattern       string `json:"pattern"`
	Type          string `json:"type"`
	Precedence    int    `json:"precedence,omitempty"`
	Associativity string `json:"associativity,omitempty"`
}

type TokenConfig struct {
	Version string     `json:"version"`
	Tokens  []TokenDef `json:"tokens"`
}

func (t TokenConfig) QueryAll(literal string) (TokenDef, error, bool) {
	for _, def := range t.Tokens {
		reg, err := regexp.Compile(def.Pattern)
		if err != nil {
			return TokenDef{}, err, true
		}

		if reg.Match([]byte(literal)) {
			return def, nil, true
		}
	}
	return TokenDef{}, nil, false
}

func LoadTokenConfig(path string) (*TokenConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config TokenConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
