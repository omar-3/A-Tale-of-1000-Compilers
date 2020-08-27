package craftinginterpreters.lox;

import java.util.List;

abstract class Expr {
	static class Binary extends Expr {
		Binary ( Expr left, Token operator, Expr right) {
			this.Expr = Expr;
			this.operator = operator;
			this.right = right;
		}
		final  Expr left;
		final Token operator;
		final Expr right;
	}
	static class Grouping extends Expr {
		Grouping ( Expr expression) {
			this.Expr = Expr;
		}
		final  Expr expression;
	}
	static class Literal extends Expr {
		Literal ( Object value) {
			this.Object = Object;
		}
		final  Object value;
	}
	static class Unary extends Expr {
		Unary ( Token operator, Expr right) {
			this.Token = Token;
			this.right = right;
		}
		final  Token operator;
		final Expr right;
	}
}
