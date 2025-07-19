package parser

import (
	. "github.com/jishaocong0910/go-sql-parser/enum"
)

func NewMySqlLexer(sql string) *mySqlLexer {
	return newMySqlLexer(sql)
}

func NewMySqlParser(sql string) *mySqlParser {
	return newMySqlParser(sql)
}

func GetCursorC(l lexer_) rune {
	return l.lexer_().cursor.c
}

func GetCursorPos(l lexer_) int {
	return l.lexer_().cursor.pos
}

func GetToken(l lexer_) Token {
	return l.token()
}

func GetTokenVal(l lexer_) string {
	return l.tokenVal()
}

func NextChar(l lexer_) rune {
	return l.lexer_().nextChar()
}

func NextToken(l lexer_) Token {
	return l.nextToken()
}

func NextTokenIncludeComment(l lexer_) Token {
	return l.nextTokenIncludeComment()
}

func HasQualifier(l *mySqlLexer) bool {
	return l.hasQualifier()
}
