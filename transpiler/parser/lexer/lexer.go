package lexer

import (
	"fmt"
	"iter"
	"strings"
	"unicode"
	"unicode/utf8" // CRITICAL: For correct Unicode handling
)

// ... (TokenType, Token, TokenResult, TokenDef, TokenConfig as defined above) ...

type Lexer struct {
	tokenConfig TokenConfig
	input       string
	// current position in input (points to the current character)
	position int
	// read position in input (points to the next character after current)
	readPosition int
	// current character under examination
	ch rune
	// current line number
	line int
	// current column number
	column int
}

func New(input string, tokenConfig TokenConfig) *Lexer {
	l := &Lexer{
		tokenConfig: tokenConfig,
		input:       input,
		line:        1,
		column:      0, // Column starts at 0 for the first char
	}
	l.readChar() // Initialize ch, position, readPosition
	return l
}

// readChar reads the next character and advances the positions.
// It handles Unicode characters correctly.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		r, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r
		l.position = l.readPosition
		l.readPosition += size

		if l.ch == '\n' {
			l.line++
			l.column = 0 // Reset column on newline
		} else {
			l.column++
		}
	}
}

// peekChar returns the character at l.readPosition without advancing.
// Useful for lookahead (e.g., checking for '==' after seeing '=').
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0 // EOF
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return r
}

// skipWhitespace consumes whitespace characters.
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

// readIdentifier reads a sequence of letters, digits, or underscores.
func (l *Lexer) readIdentifier() string {
	startPos := l.position // Store the start position of the token
	for unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// readNumber reads a sequence of digits, potentially including a decimal point.
func (l *Lexer) readNumber() string {
	startPos := l.position
	hasDecimal := false
	for unicode.IsDigit(l.ch) || (l.ch == '.' && !hasDecimal) {
		if l.ch == '.' {
			hasDecimal = true
		}
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// readString reads a string literal enclosed in double or single quotes.
func (l *Lexer) readString() (string, error) {
	quoteChar := l.ch // Store the opening quote character (' or ")
	l.readChar()      // Consume the opening quote

	var builder strings.Builder
	for l.ch != 0 && l.ch != quoteChar {
		if l.ch == '\\' { // Handle escape sequences
			l.readChar() // Consume '\'
			switch l.ch {
			case 'n':
				builder.WriteRune('\n')
			case 't':
				builder.WriteRune('\t')
			case 'r':
				builder.WriteRune('\r')
			case '\\':
				builder.WriteRune('\\')
			case quoteChar: // Escaped quote
				builder.WriteRune(quoteChar)
			default:
				// Unrecognized escape sequence, treat as literal characters
				builder.WriteRune('\\')
				builder.WriteRune(l.ch)
			}
		} else {
			builder.WriteRune(l.ch)
		}
		l.readChar()
	}

	if l.ch != quoteChar {
		return "", fmt.Errorf("unclosed string literal: missing closing '%c'", quoteChar)
	}

	l.readChar() // Consume the closing quote
	return builder.String(), nil
}

// newToken creates a Token object with current position info.
func (l *Lexer) newToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Pos:     l.position, // This needs to be the actual start of the token
		Line:    l.line,
		Column:  l.column,
	}
}

// Next provides an iterator for tokens.
func (l *Lexer) Next() iter.Seq[TokenResult] {
	return func(yield func(TokenResult) bool) {
		for l.ch != 0 { // Loop until EOF
			l.skipWhitespace() // Skip leading whitespace for each token

			if l.ch == 0 { // If current character is EOF
				yield(TokenResult{Token: l.newToken(EOF, ""), Err: nil})
				return // Terminate the sequence (iterator)
			}

			startPos := l.position // Record start position BEFORE consuming token
			startLine := l.line
			startCol := l.column

			var tok Token
			var err error

			switch {
			case unicode.IsLetter(l.ch) || l.ch == '_':
				literal := l.readIdentifier()
				if keywordType, ok := l.tokenConfig.Keywords[literal]; ok {
					tok = l.newToken(keywordType, literal)
				} else {
					tok = l.newToken(IDENTIFIER, literal)
				}

			case unicode.IsDigit(l.ch):
				literal := l.readNumber()
				if strings.Contains(literal, ".") {
					tok = l.newToken(FLOAT, literal)
				} else {
					tok = l.newToken(INTEGER, literal)
				}
			case l.ch == '"' || l.ch == '\'':
				literal, readErr := l.readString()
				if readErr != nil {
					err = readErr
					tok = l.newToken(ILLEGAL, l.input[startPos:l.readPosition])
				} else {
					tok = l.newToken(STRING, literal)
				}
			// Operators and Delimiters
			case l.ch == '=':
				if l.peekChar() == '=' {
					l.readChar()
					l.readChar()
					tok = l.newToken(EQ, "==")
				} else {
					l.readChar()
					tok = l.newToken(ASSIGN, "=")
				}
			case l.ch == '!':
				if l.peekChar() == '=' {
					l.readChar()
					l.readChar()
					tok = l.newToken(NOT_EQ, "!=")
				} else {
					l.readChar()
					tok = l.newToken(BANG, "!")
				}
			case l.ch == '<':
				if l.peekChar() == '=' {
					l.readChar()
					l.readChar()
					tok = l.newToken(LE, "<=")
				} else {
					l.readChar()
					tok = l.newToken(LT, "<")
				}
			case l.ch == '>':
				if l.peekChar() == '=' {
					l.readChar()
					l.readChar()
					tok = l.newToken(GE, ">=")
				} else {
					l.readChar()
					tok = l.newToken(GT, ">")
				}
			case l.ch == '+':
				l.readChar()
				tok = l.newToken(PLUS, "+")
			case l.ch == '-':
				l.readChar()
				tok = l.newToken(MINUS, "-")
			case l.ch == '*':
				l.readChar()
				tok = l.newToken(ASTERISK, "*")
			case l.ch == '/':
				if l.peekChar() == '/' {
					l.readChar()
					l.readChar()
					for l.ch != '\n' && l.ch != 0 {
						l.readChar()
					}
					continue
				} else if l.peekChar() == '*' {
					l.readChar()
					l.readChar()
					for {
						if l.ch == 0 {
							err = fmt.Errorf("unclosed multi-line comment")
							tok = l.newToken(ILLEGAL, l.input[startPos:l.readPosition])
							break
						}
						if l.ch == '*' && l.peekChar() == '/' {
							l.readChar()
							l.readChar()
							break
						}
						l.readChar()
					}
					if err != nil {
						break
					}
					continue
				} else {
					l.readChar()
					tok = l.newToken(SLASH, "/")
				}
			case l.ch == ';':
				l.readChar()
				tok = l.newToken(SEMICOLON, ";")
			case l.ch == ',':
				l.readChar()
				tok = l.newToken(COMMA, ",")
			case l.ch == '(':
				l.readChar()
				tok = l.newToken(LPAREN, "(")
			case l.ch == ')':
				l.readChar()
				tok = l.newToken(RPAREN, ")")
			case l.ch == '{':
				l.readChar()
				tok = l.newToken(LBRACE, "{")
			case l.ch == '}':
				l.readChar()
				tok = l.newToken(RBRACE, "}")
			case l.ch == '[':
				l.readChar()
				tok = l.newToken(LBRACKET, "[")
			case l.ch == ']':
				l.readChar()
				tok = l.newToken(RBRACKET, "]")
			default:
				// Unrecognized character
				literal := string(l.ch)
				err = fmt.Errorf("unrecognized character: '%c'", l.ch)
				tok = l.newToken(ILLEGAL, literal)
				l.readChar()
			}

			// Set correct position for the token
			tok.Pos = startPos
			tok.Line = startLine
			tok.Column = startCol

			if !yield(TokenResult{Token: tok, Err: err}) {
				return
			}
		}

		// yield an EOF token at the end
		yield(TokenResult{Token: l.newToken(EOF, ""), Err: nil})
	}
}
