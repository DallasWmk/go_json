package main

import (
	"fmt"
	"os"

	"github.com/DallasWmk/go_json/internal/lexer"
)

func main() {
	validName := "test.json"
	validFile, err := os.Open(validName)
	if err != nil {
		panic(err)
	}
	parse(validFile, validName)

	invalidName := "invalidtest.json"
	invalidFile, err := os.Open(invalidName)
	if err != nil {
		panic(err)
	}
	parse(invalidFile, invalidName)
}

func parse(file *os.File, fileName string) {
	myLexer := lexer.NewLexer(file)
	validityCheck := map[string]bool{
		"curly":   true,
		"bracket": true,
	}
	for {
		_, tok, _ := myLexer.Lex()
		if tok == lexer.EOF {
			break
		}
		switch tok {
		case lexer.CurlyOpen, lexer.CurlyClose:
			validityCheck["curly"] = !validityCheck["curly"]
		case lexer.BracketOpen, lexer.BracketClose:
			validityCheck["bracket"] = !validityCheck["bracket"]
		default:
			continue
		}
	}
	for key, value := range validityCheck {
		if !value {
			fmt.Printf("Missing token: %s in %s\n", key, fileName)
			return
		}
	}
	fmt.Printf("All tokens present in %s\n", fileName)
}
