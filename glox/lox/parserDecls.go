package lox

func (p *Parser) declaration() Stmt {
	switch {
	case p.match(TokenTypeVar):
		return p.varDeclaration()
	
	default:
		return p.statement()
	}
}

func (p *Parser) varDeclaration() Stmt {
	name := p.consume(TokenTypeIdentifier, "Expect variable name")
	
	var initializer Expr
	if p.match(TokenTypeEqual) {
		initializer = p.expression()
	}
	
	p.consume(TokenTypeSemicolon, "Expect ';' after variable declaration")
	return &StmtVar{name, initializer}
}
