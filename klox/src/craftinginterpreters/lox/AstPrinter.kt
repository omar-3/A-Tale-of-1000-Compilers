package craftinginterpreters.lox

import java.lang.StringBuilder

private object AstPrinter : Expr.Visitor<String> {

    @JvmStatic
    fun main(args: Array<String>) {
        val expression = Expr.Binary(
            Expr.Unary(
                Token(TokenType.MINUS, "-", null, 1),
                Expr.Literal(123)
            ),
            Token(TokenType.STAR, "*", null, 1),
            Expr.Grouping(Expr.Literal(45.67))
        )

        print(AstPrinter.print(expression))
    }

    private fun print(expr : Expr) : String = expr.accept(this)

    override fun visitBinaryExpr(expr: Expr.Binary): String =
        parenthesize(expr.operator.lexeme, expr.left, expr.right)

    override fun visitGroupingExpr(expr: Expr.Grouping): String =
        parenthesize("group", expr.expression)

    override fun visitLiteralExpr(expr: Expr.Literal): String =
        expr.value.toString()

    override fun visitUnaryExpr(expr: Expr.Unary): String =
        parenthesize(expr.operator.lexeme, expr.right)

    private fun parenthesize(name: String, vararg exprs : Expr) : String{
        val builder = StringBuilder()

        builder.append("(").append(name)
        for (expr in exprs) {
            builder.append(" ")
            builder.append(expr.accept(this))
        }
        builder.append(")")

        return builder.toString()
    }
}