package lexer

type TokenType string

const (
	EOF     TokenType = "EOF"     // End Of File
	ILLEGAL TokenType = "ILLEGAL" // Unrecognized token

	// Literals
	STATEMENT  TokenType = "STATEMENT"
	IDENTIFIER TokenType = "IDENTIFIER"
	INTEGER    TokenType = "INTEGER" // Changed from DIGIT to reflect numbers
	FLOAT      TokenType = "FLOAT"   // Added for floating-point numbers
	STRING     TokenType = "STRING"

	// Operators
	PLUS     TokenType = "PLUS"     // +
	MINUS    TokenType = "MINUS"    // -
	ASTERISK TokenType = "ASTERISK" // *
	SLASH    TokenType = "SLASH"    // /
	ASSIGN   TokenType = "ASSIGN"   // =
	EQ       TokenType = "EQ"       // ==
	NOT_EQ   TokenType = "NOT_EQ"   // !=
	LT       TokenType = "LT"       // <
	LE       TokenType = "LE"       // <=
	GT       TokenType = "GT"       // >
	GE       TokenType = "GE"       // >=
	BANG     TokenType = "BANG"     // !

	// Delimiters
	COMMA     TokenType = "COMMA"     // ,
	SEMICOLON TokenType = "SEMICOLON" // ;
	LPAREN    TokenType = "LPAREN"    // (
	RPAREN    TokenType = "RPAREN"    // )
	LBRACE    TokenType = "LBRACE"    // {
	RBRACE    TokenType = "RBRACE"    // }
	LBRACKET  TokenType = "LBRACKET"  // [
	RBRACKET  TokenType = "RBRACKET"  // ]
)

type Token struct {
	Type    TokenType
	Literal string
	Pos     int // Starting position of the token in the input string (byte offset)
	Line    int // Line number (for error reporting)
	Column  int // Column number (for error reporting)
}

type TokenResult struct {
	Token Token
	Err   error
}

type TokenDef struct {
	Id     string
	Regexp string
}

type TokenConfig struct {
	Keywords  map[string]TokenType
	Operators map[string]TokenType
}

func NewTokenConfig() TokenConfig {
	return TokenConfig{
		Keywords: map[string]TokenType{
			"func":    "FUNC",
			"let":     "LET",
			"true":    "TRUE",
			"false":   "FALSE",
			"if":      "IF",
			"else":    "ELSE",
			"return":  "RETURN",
			"package": "PACKAGE",
			"var":     "VAR",
			// more...
		},
		Operators: map[string]TokenType{
			"==": EQ,
			"!=": NOT_EQ,
			"<=": LE,
			">=": GE,
			// more...
		},
	}
}

func (tc TokenConfig) QueryAll(literal string) (TokenDef, error, bool) {
	if tokType, ok := tc.Keywords[literal]; ok {
		return TokenDef{Id: string(tokType)}, nil, true
	}
	return TokenDef{}, nil, false
}
