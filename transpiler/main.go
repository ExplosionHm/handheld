package main

import (
	"gopascal/parser"
	"gopascal/parser/lexer"
	"log"
)

func main() {

	code := `package main
	
	var hi int = 0

	func main() {
		println("Hello world! func")
	}
	`

	pascal, err := parser.TranspileToPascal(code, lexer.NewTokenConfig())
	if err != nil {
		log.Fatal(err)
	}
	/* var out string = ""
	for _, t := range token {
		out += "," + fmt.Sprintf("%v", t)
	} */
	if pascal.Lpr != "" {
		log.Println(pascal)
	}
}
