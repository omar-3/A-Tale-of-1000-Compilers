package lox

func (p *Parser) expression() Expr {
	return p.assignment()
}

func (p *Parser) assignment() Expr {
	expr := p.equality()
	
	if p.match(TokenTypeEqual) {
		equals := p.previous()
		value := p.assignment()
		
		exprVar, ok := value.(*ExprVar)
		if ok {
			return &ExprAssign{exprVar.Name, value}
		}
		
		panic(&ParseError{equals, "Invalid assignment target"})
	}
	
	return expr
}

func (p *Parser) equality() Expr {
	expr := p.comparison()
	
	for p.match(TokenTypeBangEqual, TokenTypeEqualEqual) {
		op := p.previous()
		right := p.comparison()
		
		expr = &ExprBinary{expr, op, right}
	}
	
	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()
	
	for p.match(TokenTypeGreater, TokenTypeGreaterEqual, TokenTypeLess, TokenTypeLessEqual) {
		op := p.previous()
		right := p.term()
		
		expr = &ExprBinary{expr, op, right}
	}
	
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()
	
	for p.match(TokenTypeMinus, TokenTypePlus) {
		op := p.previous()
		right := p.factor()
		
		expr = &ExprBinary{expr, op, right}
	}
	
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()
	
	for p.match(TokenTypeSlash, TokenTypeStar) {
		op := p.previous()
		right := p.unary()
		
		expr = &ExprBinary{expr, op, right}
	}
	
	return expr
}

func (p *Parser) unary() Expr {
	if p.match(TokenTypeBang, TokenTypeMinus) {
		op := p.previous()
		right := p.unary()
		
		return &ExprUnary{op, right}
	}
	
	return p.primary()
}

func (p *Parser) primary() Expr {
	switch {
	case p.match(TokenTypeFalse):
		return &ExprLiteral{false}
	
	case p.match(TokenTypeTrue):
		return &ExprLiteral{true}
	
	case p.match(TokenTypeNil):
		return &ExprLiteral{nil}
	
	case p.match(TokenTypeNumber, TokenTypeString):
		return &ExprLiteral{p.previous().Literal}
	
	case p.match(TokenTypeIdentifier):
		return &ExprVar{p.previous()}
	
	case p.match(TokenTypeLeftParen):
		expr := p.expression()
		
		p.consume(TokenTypeRightParen, "Expect ')' after expression.")
		return &ExprGrouping{expr}
	
	default:
		panic(&ParseError{p.peek(), "Expect expression"})
	}
}

