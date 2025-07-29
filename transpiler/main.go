package main

import "gopascal/parser"

func main() {

	code := `package main
	
	var hi int = 0

	func main() {
		println("Hello world! func")
	}
	`

	parser.TranspileToPascal(code)
}
