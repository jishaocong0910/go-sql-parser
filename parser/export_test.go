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

func GetCursorC(l i_Lexer) rune {
	return l.m_EC05053E2C60().cursor.c
}

func GetCursorPos(l i_Lexer) int {
	return l.m_EC05053E2C60().cursor.pos
}

func GetToken(l i_Lexer) Token {
	return l.token()
}

func GetTokenVal(l i_Lexer) string {
	return l.tokenVal()
}

func NextChar(l i_Lexer) rune {
	return l.m_EC05053E2C60().nextChar()
}

func NextToken(l i_Lexer) Token {
	return l.nextToken()
}

func NextTokenIncludeComment(l i_Lexer) Token {
	return l.nextTokenIncludeComment()
}

func HasQualifier(l *mySqlLexer) bool {
	return l.hasQualifier()
}
