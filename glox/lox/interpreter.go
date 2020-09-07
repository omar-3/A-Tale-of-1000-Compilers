package lox

import (
	"fmt"
)

type Interpreter struct {
	environment *Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{NewEnvironment(nil)}
}

func (i *Interpreter) Interpret(stmts []Stmt) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	
	for _, s := range stmts {
		i.execute(s)
	}
	
	return nil
}

func (i *Interpreter) execute(stmt Stmt) {
	switch t := stmt.(type) {
	case *StmtExpr:
		i.stmtExpr(t)
	
	case *StmtBlock:
		i.stmtBlock(t)
	
	case *StmtPrint:
		i.stmtPrint(t)
	
	case *StmtVar:
		i.stmtVar(t)
	}
}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	switch t := expr.(type) {
	case *ExprAssign:
		return i.exprAssign(t)
	
	case *ExprBinary:
		return i.exprBinary(t)
	
	case *ExprGrouping:
		return i.exprGrouping(t)
	
	case *ExprLiteral:
		return i.exprLiteral(t)
	
	case *ExprUnary:
		return i.exprUnary(t)
	
	case *ExprVar:
		return i.exprVar(t)
	
	default:
		return nil
	}
}

// expressions

func (i *Interpreter) exprAssign(e *ExprAssign) interface{} {
	val := i.evaluate(e.Value)
	i.environment.Assign(e.Name.Lexeme, val)
	return val
}

func (i *Interpreter) exprBinary(e *ExprBinary) interface{} {
	left := i.evaluate(e.Left)
	right := i.evaluate(e.Right)
	
	switch e.Operator.Type {
	case TokenTypeBangEqual:
		return !equal(left, right)
	
	case TokenTypeEqualEqual:
		return equal(left, right)
	
	case TokenTypeGreater:
		l, r := checkNumberOperands(e.Operator, left, right)
		return l > r
	
	case TokenTypeGreaterEqual:
		l, r := checkNumberOperands(e.Operator, left, right)
		return l >= r
	
	case TokenTypeLess:
		l, r := checkNumberOperands(e.Operator, left, right)
		return l < r
	
	case TokenTypeLessEqual:
		l, r := checkNumberOperands(e.Operator, left, right)
		return l <= r
	
	case TokenTypePlus:
		// if the left operand is a string, we stringify the right operand,
		// and then concatenate.
		// if the left operand is a number, the right operand must be a number
		
		str, isStr := left.(StringLiteral)
		if isStr {
			return string(str) + fmt.Sprint(right)
		} else {
			l, r := checkNumberOperands(e.Operator, left, right)
			return l + r
		}
	
	case TokenTypeMinus:
		l, r := checkNumberOperands(e.Operator, left, right)
		return l - r
	
	case TokenTypeStar:
		l, r := checkNumberOperands(e.Operator, left, right)
		return l * r
	
	case TokenTypeSlash:
		l, r := checkNumberOperands(e.Operator, left, right)
		return l / r
	}
	
	return nil
}

func (i *Interpreter) exprGrouping(e *ExprGrouping) interface{} {
	return i.evaluate(e.Expr)
}

func (i *Interpreter) exprLiteral(e *ExprLiteral) interface{} {
	return e.Value
}

func (i *Interpreter) exprUnary(e *ExprUnary) interface{} {
	right := i.evaluate(e.Right)
	
	switch e.Operator.Type {
	case TokenTypeBang:
		return !truthy(right)
		
	case TokenTypeMinus:
		return -checkNumberOperand(e.Operator, e.Right)
	
	default:
		return nil
	}
}

func (i *Interpreter) exprVar(e *ExprVar) interface{} {
	return i.environment.Lookup(e.Name.Lexeme)
}

// statements

func (i *Interpreter) stmtExpr(stmt *StmtExpr) {
	i.evaluate(stmt.Expr)
}

func (i *Interpreter) stmtBlock(stmt *StmtBlock) {
	previous := i.environment
	
	i.environment = NewEnvironment(previous)
	defer func() { i.environment = previous }()
	
	for _,s := range stmt.stmts {
		i.execute(s)
	}
}

func (i *Interpreter) stmtPrint(stmt *StmtPrint) {
	val := i.evaluate(stmt.Expr)
	fmt.Printf("%v\n", val)
}

func (i *Interpreter) stmtVar(stmt *StmtVar) {
	var value interface{}
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}
	
	i.environment.Define(stmt.Name.Lexeme, value)
}

// helpers

func truthy(i interface{}) bool {
	switch val := i.(type) {
	case nil:
		return false
	
	case bool:
		return val
	
	default:
		return true
	}
}

func equal(left, right interface{}) bool {
	if left == nil {
		return right == nil
	} else {
		return left == right
	}
}

func checkNumberOperand(op Token, num interface{}) NumberLiteral {
	n, ok := num.(NumberLiteral)
	
	if !ok {
		panic(&RuntimeError{op, "Operand must be a number"})
	}
	
	return n
}

func checkNumberOperands(op Token, left, right interface{}) (NumberLiteral, NumberLiteral) {
	l, lok := left.(NumberLiteral)
	r, rok := right.(NumberLiteral)
	
	if !lok || !rok {
		panic(&RuntimeError{op, "Operands must both be numbers"})
	}
	
	return l, r
}
