package lox

import "fmt"

type RuntimeError struct {
	token Token
	msg string
}

func (re *RuntimeError) Error() string {
	return fmt.Sprintf("line #%d at '%v': '%s'", re.token.Line, re.token.Lexeme, re.msg)
}

type ParseError RuntimeError

func (pe *ParseError) Error() string {
	if pe.token.Type == TokenTypeEOF {
		return fmt.Sprintf("line #%d at end: %s", pe.token.Line, pe.msg)
	} else {
		return fmt.Sprintf("line #%d at '%v': %s", pe.token.Line, pe.token.Lexeme, pe.msg)
	}
}

type UndefinedVariableError struct {
	name string
}

func (e *UndefinedVariableError) Error() string {
	return fmt.Sprintf("Undefined variable '%s'", e.name)
}
