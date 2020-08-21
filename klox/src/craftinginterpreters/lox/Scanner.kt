package craftinginterpreters.lox

import java.util.ArrayList
import java.util.HashMap
import kotlin.collections.List
import kotlin.collections.Map

import craftinginterpreters.lox.TokenType.*

class Scanner(val source : String) {
    private val tokens = ArrayList<Token>()
    private var start : Int = 0
    private var current : Int= 0
    private var line : Int = 1


    private val keywords : Map<String, String> = mapOf(
        "and" to "AND",
        "class" to "CLASS",
        "else" to "ELSE",
        "false" to "FALSE",
        "for" to "FOR",
        "fun" to "FUN",
        "if" to "IF",
        "nil" to "NIL",
        "or" to "OR",
        "print" to "PRINT",
        "return" to "RETURN",
        "super" to "SUPER",
        "this" to "THIS",
        "true" to "TRUE",
        "var" to "VAR",
        "while" to "WHILE",
        "break" to "BREAK"
    )

    fun scanTokens() : List<Token> {
        while (!isAtEnd) {
            start = current;
            scanToken()
        }
        tokens.add(Token(EOF, "", null, line))
        return tokens
    }

    private fun scanToken() {
        when (val c = advance()) {
            '(' -> addToken(LEFT_PAREN)
            ')' -> addToken(RIGHT_PAREN)
            '{' -> addToken(LEFT_BRACE)
            '}' -> addToken(RIGHT_BRACE)
            ',' -> addToken(COMMA)
            '.' -> addToken(DOT)
            '-' -> addToken(MINUS)
            '+' -> addToken(PLUS)
            ';' -> addToken(SEMICOLON)
            '*' -> addToken(STAR)
            '!' -> addToken(if (match('=')) BANG_EQUAL else BANG)
            '=' -> addToken(if (match('=')) EQUAL_EQUAL else EQUAL)
            '<' -> addToken(if (match('=')) LESS_EQUAL else LESS)
            '>' -> addToken(if (match('=')) GREATER_EQUAL else GREATER)
            '/' -> when {
                match('/') ->
                    while(peek != '\n' && !isAtEnd) advance()
                else ->
                    addToken(SLASH)
            }
            ' ', '\r', '\t' -> {}
            '\n' -> line++
            '"' -> string()
            else -> when {
                isDigit(c) -> {
                    number()
                }
                isAlpha(c) -> {
                    identifier()
                }
                else -> Lox.error(line, "Unexpected character")
            }
        }

    }

    private fun identifier() {
        while (isAlphaNumeric(peek)) advance()
        val text = source.substring(start, current)
        var type : TokenType = TokenType.valueOf(keywords[text]?: "IDENTIFIER")
        addToken(type)
    }

    private fun number() {
        while (isDigit(peek)) advance()
        if (peek == '.' && isDigit(peekNext)) {
            advance()
            while (isDigit(peek)) advance()
        }
        addToken(NUMBER, (source.substring(start, current)).toDouble())
    }

    private fun string() {
        while (peek != '"' && !isAtEnd) {
            if (peek == '\n') line++;
            advance()
        }

        if (isAtEnd) {
            Lox.error(line, "Unterminated string.")
            return
        }

        advance()
        val value = source.substring(start+1, current -1)
        addToken(STRING, value)
    }

    private val isAtEnd : Boolean
        get() = current >= source.length

    private val peek : Char
        get() = if (isAtEnd) '\u0000' else source[current]

    private val peekNext : Char
        get() = if (current + 1 >= source.length) '\u0000' else source[current + 1]

    private fun isDigit(c : Char) : Boolean { return c in '0'..'9' }
    private fun isAlpha(c : Char) : Boolean {
        return (c in 'a'..'z') || (c in 'A'..'Z') || c == '_'
    }
    private fun isAlphaNumeric(c : Char) : Boolean { return isDigit(c) || isAlpha(c) }
    private fun advance() : Char {
        current++
        return source[current - 1]
    }

    private fun addToken(type : TokenType, literal: Any? = null) {
        val text = source.substring(start, current)
        tokens.add(Token(type, text, literal, line))
    }

    private fun match(expected : Char) : Boolean {
        if (isAtEnd) return false
        if (source[current] != expected) return false

        current++
        return true
    }
}