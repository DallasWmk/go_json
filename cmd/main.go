package main

import (
	"fmt"
	"os"

	"github.com/DallasWmk/go_json/internal/lexer"
)

func main() {
	file, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	myLexer := lexer.NewLexer(file)
	myMap := make(map[string]int)
	for {
		_, tok, _ := myLexer.Lex()
		if tok == lexer.EOF {
			break
		}
		myMap[tok.String()]++
	}
	if checkValidity(myMap) {
		fmt.Println("JSON is valid")
	} else {
		fmt.Println("JSON is invalid")
	}
}

func checkValidity(myMap map[string]int) bool {
	if myMap[lexer.Tokens[lexer.CurlyOpen]] != myMap[lexer.Tokens[lexer.CurlyClose]] {
		return false
	}
	if myMap[lexer.Tokens[lexer.BracketOpen]] != myMap[lexer.Tokens[lexer.BracketClose]] {
		return false
	}
	return true
}
