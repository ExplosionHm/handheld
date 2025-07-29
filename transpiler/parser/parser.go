package parser

import (
	"gopascal/parser/lexer"
	"log"
)

func TranspileToPascal(code string) (string, error) {
	tokenDef, err := lexer.LoadTokenConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	l := lexer.New(code, *tokenDef)

	for result := range l.Next() {
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		if !result.IsZero {
			log.Println(result.Token)
		}
	}

	return "", nil
}
