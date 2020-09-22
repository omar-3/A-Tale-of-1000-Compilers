package lox

import (
	"fmt"
)

// Token is the object that what
// the scanner will emit to be
// consumed by the parser
type Token struct {
	Type    TokenType
	Lexeme  string
	Literal fmt.Stringer
	Line    int
}

// GoString is implementation for the interface
// fmt.Stringer that handels how the struct is displayed
// when struct instant is a placeholder for %#v in formatted
// String
func (t Token) GoString() string {
	return fmt.Sprintf("%s %s %#v", t.Type, t.Lexeme, t.Literal)
}
