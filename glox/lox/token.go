package lox

import "fmt"

type TokenType uint8

//go:generate stringer -type=TokenType
const (
	TokenTypeUnknown TokenType = iota
	
	// single character tokens
	TokenTypeLeftParen
	TokenTypeRightParen
	TokenTypeLeftBrace
	TokenTypeRightBrace
	TokenTypeComma
	TokenTypeDot
	TokenTypeMinus
	TokenTypePlus
	TokenTypeSemicolon
	TokenTypeSlash
	TokenTypeStar
	
	// one or two character tokens
	TokenTypeBang
	TokenTypeBangEqual
	TokenTypeEqual
	TokenTypeEqualEqual
	TokenTypeGreater
	TokenTypeGreaterEqual
	TokenTypeLess
	TokenTypeLessEqual
	
	// literals
	TokenTypeIdentifier
	TokenTypeString
	TokenTypeNumber
	
	// reservedWords
	TokenTypeAnd
	TokenTypeClass
	TokenTypeElse
	TokenTypeFalse
	TokenTypeFun
	TokenTypeFor
	TokenTypeIf
	TokenTypeNil
	TokenTypeOr
	TokenTypePrint
	TokenTypeReturn
	TokenTypeSuper
	TokenTypeThis
	TokenTypeTrue
	TokenTypeVar
	TokenTypeWhile
	
	// extra
	TokenTypeEOF
)

var reservedWords = map[string]TokenType{
	"and":    TokenTypeAnd,
	"class":  TokenTypeClass,
	"else":   TokenTypeElse,
	"false":  TokenTypeFalse,
	"fun":    TokenTypeFun,
	"for":    TokenTypeFor,
	"if":     TokenTypeIf,
	"nil":    TokenTypeNil,
	"or":     TokenTypeOr,
	"print":  TokenTypePrint,
	"return": TokenTypeReturn,
	"super":  TokenTypeSuper,
	"this":   TokenTypeThis,
	"true":   TokenTypeTrue,
	"var":    TokenTypeVar,
	"while":  TokenTypeWhile,
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal fmt.Stringer
	Line    int
}

func (tok *Token) String() string {
	return fmt.Sprintf("Token{ Type: '%s', Lexeme: '%s', Literal: '%s', Line: %d }", tok.Type, tok.Lexeme, tok.Literal, tok.Line)
}
