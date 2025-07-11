package lexer

import (
	"bufio"
	"strings"
	"testing"
)

func TestLexBoolTrue(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("true"))
	myLexer := NewLexer(reader)
	ret := myLexer.lexBool()
	if ret != "true" {
		t.Errorf("Expected 'true', got '%s'", ret)
	}
}

func TestLexBoolFalse(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("false"))
	myLexer := NewLexer(reader)
	ret := myLexer.lexBool()
	if ret != "false" {
		t.Errorf("Expected 'false', got '%s'", ret)
	}
}

func TestLexBoolInvalid(t *testing.T) {
	// string with a non-boolean value
	reader := bufio.NewReader(strings.NewReader("test"))
	myLexer := NewLexer(reader)
	ret := myLexer.lexBool()
	if ret != "" {
		t.Errorf("Expected empty string, got '%s'", ret)
	}
}

func TestPeek(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("foobar"))
	myLexer := NewLexer(reader)
	ret := myLexer.peek(3)
	if ret != "foo" {
		t.Errorf("Expected foo, got '%s'", ret)
	}
}

func TestLexIntValid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("1234"))
	myLexer := NewLexer(reader)
	ret := myLexer.lexInt()
	if ret != "1234" {
		t.Errorf("Expected 1234, got '%s'", ret)
	}
}

func TestLexIntInvalid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("test"))
	myLexer := NewLexer(reader)
	ret := myLexer.lexInt()
	if ret != "" {
		t.Errorf("Expected empty string, got '%s'", ret)
	}
}

func TestLexDblQuoteInvalid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`"test`))
	myLexer := NewLexer(reader)
	ret := myLexer.lexQuote()
	if ret != "" {
		t.Errorf("Expected empty string, got '%s'", ret)
	}
}

func TestLexDblQuoteValid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`"test"`))
	myLexer := NewLexer(reader)
	ret := myLexer.lexQuote()
	if ret != `"test"` {
		t.Errorf(`Expected "test", got '%s'`, ret)
	}
}

func TestLexSingleQuoteInvalid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`'test`))
	myLexer := NewLexer(reader)
	ret := myLexer.lexQuote()
	if ret != "" {
		t.Errorf("Expected empty string, got '%s'", ret)
	}
}

func TestLexSingleQuoteValid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`'test'`))
	myLexer := NewLexer(reader)
	ret := myLexer.lexQuote()
	if ret != `'test'` {
		t.Errorf(`Expected 'test', got '%s'`, ret)
	}
}
