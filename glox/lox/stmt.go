package lox

type Stmt interface {
	// essentially a no-op function simply to distinguish Stmts from generic interface{}
	Statement() Stmt
}

type StmtExpr struct {
	Expr
}

type StmtBlock struct {
	stmts []Stmt
}

type StmtPrint struct {
	Expr
}

type StmtVar struct {
	Name Token
	Initializer Expr
}

func (s *StmtExpr) Statement() Stmt { return s }

func (s *StmtBlock) Statement() Stmt { return s }

func (s *StmtPrint) Statement() Stmt { return s }

func (s *StmtVar) Statement() Stmt { return s }
