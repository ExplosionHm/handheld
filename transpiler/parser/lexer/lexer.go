package lexer

import (
	"iter"
	"log"
	"regexp"
	"unicode"
)

type Lexer struct {
	tokenConfig TokenConfig
	input       string
	index       int
	maxIndex    int
}

func New(input string, tokenConfig TokenConfig) *Lexer {
	return &Lexer{
		tokenConfig: tokenConfig,
		input:       input,
		maxIndex:    len(input) - 1,
	}
}

func (l *Lexer) advance() {
	if l.index < len(l.input) {
		l.index++
	}
}

func (l *Lexer) readIdentifier() string {
	for l.index < len(l.input) && unicode.IsSpace(rune(l.input[l.index])) {
		l.advance()
	}

	start := l.index

	for l.index < len(l.input) {
		r := rune(l.input[l.index])
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			l.advance()
		} else {
			break
		}
	}

	return l.input[start:l.index]
}

func (l *Lexer) readDigit() string {
	for l.index < len(l.input) && unicode.IsSpace(rune(l.input[l.index])) {
		l.advance()
	}

	start := l.index
	for l.index < len(l.input) {
		r := rune(l.input[l.index])
		if unicode.IsDigit(r) || r == '_' {
			l.advance()
		} else {
			break
		}
	}
	return l.input[start:l.index]
}

func (l *Lexer) readString() string {
	for l.index < len(l.input) && unicode.IsSpace(rune(l.input[l.index])) {
		l.advance()
	}

	quote, _ := regexp.Compile("[\"']")
	l.advance()
	start := l.index
	for quote.Match([]byte{l.input[l.index]}) && l.index < len(l.input) {
		l.advance()
	}
	value := l.input[start:l.index]
	l.advance()
	return value
}

func (l *Lexer) Next() iter.Seq[TokenResult] {
	return func(yield func(TokenResult) bool) {
		for l.index < l.maxIndex {
			for l.index < len(l.input) && unicode.IsSpace(rune(l.input[l.index])) {
				l.advance()
			}

			if l.index >= len(l.input) {
				break
			}

			c := rune(l.input[l.index])

			if unicode.IsLetter(c) {
				literal := l.readIdentifier()

				if def, err, ok := l.tokenConfig.QueryAll(literal); ok {
					if err != nil {
						log.Fatal(err)
					}
					tok := Token{
						Type:    TokenType(def.Id),
						Literal: literal,
						Pos:     l.index,
					}

					result := ResultToken(tok, nil)
					if !yield(result) {
						break
					}
				}
				continue
			}

			if unicode.IsDigit(c) {
				literal := l.readDigit()

				tok := Token{
					Type:    TokenType("DIGIT"),
					Literal: literal,
					Pos:     l.index,
				}

				result := ResultToken(tok, nil)
				if !yield(result) {
					break
				}
				continue
			}

			if unicode.IsSymbol(c) {
				literal := l.readString()

				tok := Token{
					Type:    TokenType("STRING"),
					Literal: literal,
					Pos:     l.index,
				}
				result := ResultToken(tok, nil)
				if !yield(result) {
					break
				}
				continue
			}

			l.advance()
		}
	}
}
