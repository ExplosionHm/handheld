package parser

import "gopascal/parser/lexer"

type Pascal struct {
	Lpi *LazarusInformation
	Lpr string
}

func TranslateToken(tok lexer.Token) string {
	return ""
}

func lookupToken(tok lexer.Token) string {
	return ""
}
