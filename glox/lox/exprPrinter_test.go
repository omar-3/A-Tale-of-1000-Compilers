package lox

import (
	"fmt"
	"go/printer"
)

func ExampleExpressionPrinter_Print() {
	expr := ExprBinary{
		Left: &ExprUnary{
			Operator: Token{TokenTypeMinus, "-", nil, 1},
			Right: &ExprLiteral{123},
		},
		Operator: Token{TokenTypeStar, "*", nil, 1},
		Right: &ExprGrouping{
			Expr: &ExprLiteral{45.67},
		},
	}
	
	fmt.Println(PrintExpr(&expr))
	// Output:
	// (* (- 123) (group 45.67))
}
