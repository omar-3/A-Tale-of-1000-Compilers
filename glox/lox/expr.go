package lox

type Expr interface {
	// essentially a no-op function simply to distinguish Exprs from generic interface{}
	Expression() Expr
}

type ExprAssign struct {
	Name Token
	Value Expr
}

type ExprBinary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

type ExprGrouping struct {
	Expr
}

type ExprLiteral struct {
	Value interface{}
}

type ExprUnary struct {
	Operator Token
	Right    Expr
}

type ExprVar struct {
	Name Token
}

func (e *ExprAssign) Expression() Expr { return e }

func (e *ExprBinary) Expression() Expr { return e }

func (e *ExprGrouping) Expression() Expr { return e }

func (e *ExprLiteral) Expression() Expr { return e }

func (e *ExprUnary) Expression() Expr { return e }

func (e *ExprVar) Expression() Expr { return e }
