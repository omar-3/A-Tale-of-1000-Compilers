package craftinginterpreters.lox.tool

import java.io.IOException
import java.io.PrintWriter
import java.util.Arrays

object GenerateAst {
    @Throws (IOException::class)
    @JvmStatic
    fun main(args: Array<String>) {
        if (args.size != 1) {
            println("Usage: generate_ast <output directory>")
            System.exit(64)
        }

        val outputDir = args[0]

        defineAst(
            outputDir, "Expr", Arrays.asList(
                "Binary   : Expr left, Token operator, Expr right",
                "Grouping : Expr expression",
                "Literal  : Object value",
                "Unary    : Token operator, Expr right"
            )
        )
    }

    @Throws (IOException::class)
    private fun defineAst(outputDir: String, baseName: String, types: List<String>) {
        val path = "$outputDir/$baseName.java"
        val writer = PrintWriter(path, "UTF-8")

        writer.println("package craftinginterpreters.lox")
        writer.println()
        writer.println("import java.util.List")
        writer.println()
        writer.println("abstract class $baseName {")

        for (type in types) {
            val className = type.split(":").map {string -> string.trim()}[0]
            val fields = type.split(":").map {string -> string.trim()}[1]
            defineType(writer, baseName, className, fields)
        }

        writer.println("}")
        writer.close()
    }

    @Throws (IOException::class)
    private fun defineType(writer: PrintWriter, baseName: String, className: String, fieldList: String) {
        writer.println("\tstatic class $className extends $baseName {")
        writer.println("\t\t$className($fieldList) {")
        val fields = fieldList.split(",").map {char -> char.trim()}
        for (field in fields) {
            val name = field.split(" ")
            writer.println("\t\t\tthis.$name = this.$name")
        }
        writer.println("\t\t}")
        writer.println()

        for (field in fields) {
            writer.println("\t\tfinal $field")
        }
        writer.println("\t}")
    }
}