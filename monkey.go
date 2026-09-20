package main

import (
	"fmt"
	"monkey/lexer"
	"monkey/token"
)

func main() {
	input := "=+(){},;"

	l := lexer.New(input)

	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)

		if tok.Type == token.EOF {
			break
		}
	}
}
