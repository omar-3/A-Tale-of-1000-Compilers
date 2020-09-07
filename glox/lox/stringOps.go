package lox

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlpha(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

func isAlphaNumeric(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

func getIdentifier(str string) (Token, int) {
	length := len(str)
	idx := 0
	
	for idx < length && isAlphaNumeric(str[idx]) {
		idx++
	}
	
	tokenType := TokenTypeIdentifier
	lexeme := str[:idx]
	
	if keyword, is := reservedWords[lexeme]; is {
		tokenType = keyword
	}
	
	return Token{tokenType, lexeme, nil, 0}, idx
}
