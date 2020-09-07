package lox

func (p *Parser) statement() Stmt {
	switch {
	case p.match(TokenTypeLeftBrace):
		return &StmtBlock{p.blockStmt()}
	
	case p.match(TokenTypePrint):
		return p.printStmt()
	
	default:
		return p.expressionStmt()
	}
}

func (p *Parser) blockStmt() []Stmt {
	stmts := []Stmt{}
	
	for !p.check(TokenTypeRightBrace) && !p.end() {
		stmts = append(stmts, p.declaration())
	}
	
	p.consume(TokenTypeRightBrace, "Expect '}' at end of block")
	
	return stmts
}

func (p *Parser) printStmt() Stmt {
	expr := p.expression()
	p.consume(TokenTypeSemicolon, "Expect ';' after value")
	
	return &StmtPrint{expr}
}

func (p *Parser) expressionStmt() Stmt {
	expr := p.expression()
	p.consume(TokenTypeSemicolon, "Expect ';' after expression")
	
	return &StmtExpr{expr}
}
