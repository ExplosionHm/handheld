package parser

import (
	"gopascal/parser/lexer"
	"log"
)

func Tokenize(code string, config lexer.TokenConfig) ([]lexer.Token, error) {
	l := lexer.New(code, lexer.NewTokenConfig())
	results := []lexer.Token{}
	for result := range l.Next() {
		if result.Err != nil {
			log.Fatal(result.Err)
		}
		log.Println(result.Token)
		results = append(results, result.Token) //! `append` is inefficent
	}
	return results, nil
}

func TranspileToPascal(code string, config lexer.TokenConfig) (pas Pascal, err error) {
	tokens, err := Tokenize(code, config)
	if err != nil {
		return Pascal{}, err
	}
	log.Println(tokens)
	return Pascal{}, nil
}
