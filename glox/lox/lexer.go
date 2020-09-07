package lox

import (
	"fmt"
	"strconv"
	"strings"
)

type StringLiteral string
type NumberLiteral float64

func (sl StringLiteral) String() string {
	return string(sl)
}

func (nl NumberLiteral) String() string {
	return fmt.Sprint(float64(nl))
}

func Lex(str string) []Token {
	var tokens []Token

	lineNumber := 0

	length := len(str)
	current := 0

	for current < length {
		tok, next, newLines := scanToken(str[current:])

		if tok != (Token{}) {
			tok.Line = lineNumber
			tokens = append(tokens, tok)
		}

		lineNumber += newLines
		current += next
	}

	tokens = append(tokens, Token{TokenTypeEOF, "", nil, lineNumber})

	return tokens
}

func scanToken(str string) (tok Token, next int, newLines int) {
	next = 1
	newLines = 0

	switch ch := str[0]; ch {
	// whitespace
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		newLines++

	// single character tokens
	case '(':
		tok = Token{TokenTypeLeftParen, "(", nil, 0}
		
	case ')':
		tok = Token{TokenTypeRightParen, ")", nil, 0}
		
	case '{':
		tok = Token{TokenTypeLeftBrace, "{", nil, 0}
		
	case '}':
		tok = Token{TokenTypeRightBrace, "}", nil, 0}
		
	case ',':
		tok = Token{TokenTypeComma, ",", nil, 0}
		
	case '.':
		tok = Token{TokenTypeDot, ".", nil, 0}
		
	case '-':
		tok = Token{TokenTypeMinus, "-", nil, 0}
		
	case '+':
		tok = Token{TokenTypePlus, "+", nil, 0}
		
	case ';':
		tok = Token{TokenTypeSemicolon, ";", nil, 0}
		
	case '/':
		tok = Token{TokenTypeSlash, "/", nil, 0}
		
	case '*':
		tok = Token{TokenTypeStar, "*", nil, 0}
		
	case '#':
		// comments go to the end of the line!
		nl := strings.Index(str, "\n")
		next += nl
		if nl != -1 {
			newLines++
		}

	// one or two character tokens
	case '!':
		if nextIs('=', str) {
			next++
			tok = Token{TokenTypeBangEqual, str[:2], nil, 0}
		} else {
			tok = Token{TokenTypeBang, "!", nil, 0}
		}
		
	case '=':
		if nextIs('=', str) {
			next++
			tok = Token{TokenTypeEqualEqual, "==", nil, 0}
		} else {
			tok = Token{TokenTypeEqual, "=", nil, 0}
		}
		
	case '>':
		if nextIs('=', str) {
			next++
			tok = Token{TokenTypeGreaterEqual, ">=", nil, 0}
		} else {
			tok = Token{TokenTypeGreater, ">", nil, 0}
		}
		
	case '<':
		if nextIs('=', str) {
			next++
			tok = Token{TokenTypeLessEqual, "<=", nil, 0}
		} else {
			tok = Token{TokenTypeLess, "<", nil, 0}
		}

	// literals
	case '"':
		// we don't want Index to grab the first quote, but we want the full index, so add 1
		end := strings.Index(str[1:], "\"") + 1
		if end != -1 {
			newLines += strings.Count(str[1:end], "\n")
			next += end
			tok = Token{TokenTypeString, str[0 : end+1], StringLiteral(str[1:end]), 0}
		} else {
			// error
		}

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		length := len(str)
		for next < length && isDigit(str[next]) {
			next++
		}

		if next < length && str[next] == '.' {
			if next+1 < length && isDigit(str[next+1]) {
				next += 2
			}
		}

		for next < length && isDigit(str[next]) {
			next++
		}

		strVal := str[0:next]
		val, _ := strconv.ParseFloat(strVal, 64)
		tok = Token{TokenTypeNumber, strVal, NumberLiteral(val), 0}

	default:
		if isAlpha(ch) {
			tok, next = getIdentifier(str)
		} else {
			// error
		}
	}

	return
}

func nextIs(want byte, str string) bool {
	return len(str) > 0 && want == str[1]
}
