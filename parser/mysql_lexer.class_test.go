package parser_test

import (
	. "go-sql-parser/enum"
	"go-sql-parser/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMySqlLexer_NextToken(t *testing.T) {
	r := require.New(t)
	// eoi
	l := parser.NewMySqlLexer("")
	r.Equal(Tokens.EOI, parser.NextToken(l))
	// identifier
	l = parser.NewMySqlLexer("abc `abc` 中文")
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	r.Equal(false, parser.HasQualifier(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	r.Equal(true, parser.HasQualifier(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("中文", parser.GetTokenVal(l))
	r.Equal(false, parser.HasQualifier(l))
	// hexadecimalNumber
	l = parser.NewMySqlLexer("0x3f 0x4D7953514C")
	r.Equal(Tokens.HEXADECIMAL_NUMBER, parser.NextToken(l))
	r.Equal("0x3f", parser.GetTokenVal(l))
	r.Equal(Tokens.HEXADECIMAL_NUMBER, parser.NextToken(l))
	r.Equal("0x4D7953514C", parser.GetTokenVal(l))
	// binaryNumber
	l = parser.NewMySqlLexer("0b11010")
	r.Equal(Tokens.BINARY_NUMBER, parser.NextToken(l))
	r.Equal("0b11010", parser.GetTokenVal(l))
	// decimalNumber
	l = parser.NewMySqlLexer("0124 9827.893 6153e+21 8378E-21")
	r.Equal(Tokens.DECIMAL_NUMBER, parser.NextToken(l))
	r.Equal("0124", parser.GetTokenVal(l))
	r.Equal(Tokens.DECIMAL_NUMBER, parser.NextToken(l))
	r.Equal("9827.893", parser.GetTokenVal(l))
	r.Equal(Tokens.DECIMAL_NUMBER, parser.NextToken(l))
	r.Equal("6153e+21", parser.GetTokenVal(l))
	r.Equal(Tokens.DECIMAL_NUMBER, parser.NextToken(l))
	r.Equal("8378E-21", parser.GetTokenVal(l))
	r.Equal(Tokens.EOI, parser.NextToken(l))
	// comma
	l = parser.NewMySqlLexer(",")
	r.Equal(Tokens.COMMA, parser.NextToken(l))
	r.Equal(",", parser.GetTokenVal(l))
	// lParen
	l = parser.NewMySqlLexer("(")
	r.Equal(Tokens.L_PAREN, parser.NextToken(l))
	r.Equal("(", parser.GetTokenVal(l))
	// rParen
	l = parser.NewMySqlLexer(")")
	r.Equal(Tokens.R_PAREN, parser.NextToken(l))
	r.Equal(")", parser.GetTokenVal(l))
	// dot
	l = parser.NewMySqlLexer(".")
	r.Equal(Tokens.DOT, parser.NextToken(l))
	r.Equal(".", parser.GetTokenVal(l))
	// string
	l = parser.NewMySqlLexer("'abc' 'abcwef'")
	r.Equal(Tokens.STRING, parser.NextToken(l))
	r.Equal("'abc'", parser.GetTokenVal(l))
	r.Equal(Tokens.STRING, parser.NextToken(l))
	r.Equal("'abcwef'", parser.GetTokenVal(l))
	// ignoreComment
	l = parser.NewMySqlLexer("select" +
		"  name\n,-- this COMMENT continues to the end OF line \n" +
		"  #this COMMENT continues to the end OF line \n" +
		"  `age`,#this COMMENT continues to the end OF line\r\n" +
		"  -- this COMMENT continues to the end OF line \r\n" +
		"  dept" +
		"  /*this IS a\n" +
		"  multiple-line\r\n" +
		"  com*ment*/" +
		"FROM" +
		"  t_employee ;")
	parser.NextToken(l) // select
	parser.NextToken(l) // name
	parser.NextToken(l) // ,
	parser.NextToken(l) // `age`
	r.Equal("age", parser.GetTokenVal(l))
	parser.NextToken(l)
	parser.NextToken(l)
	r.Equal("dept", parser.GetTokenVal(l))
	parser.NextToken(l)
	r.Equal(Tokens.FROM, parser.GetToken(l))
	// nextComment
	l = parser.NewMySqlLexer("select" +
		"  name\n,-- this COMMENT continues to the end OF line \n" +
		"  #this COMMENT continues to the end OF line\n" +
		"  `age`,#this COMMENT continues to the end OF line\r\n" +
		"  -- this COMMENT continues to the end OF line \r\n" +
		"  dept" +
		"  /*this IS a\n" +
		"  multiple-line\r\n" +
		"  com*ment*/" +
		"FROM" +
		"  t_employee ;")
	parser.NextToken(l) // select
	parser.NextToken(l) // name
	parser.NextToken(l) // ,
	r.Equal(Tokens.COMMENT, parser.NextTokenIncludeComment(l))
	r.Equal("-- this COMMENT continues to the end OF line ", parser.GetTokenVal(l))
	r.Equal(Tokens.COMMENT, parser.NextTokenIncludeComment(l))
	r.Equal("#this COMMENT continues to the end OF line", parser.GetTokenVal(l))
	parser.NextToken(l) // `age`
	r.Equal("age", parser.GetTokenVal(l))
	parser.NextToken(l) // ,
	parser.NextToken(l)
	r.Equal("dept", parser.GetTokenVal(l))
	r.Equal(Tokens.COMMENT, parser.NextTokenIncludeComment(l))
	r.Equal("/*this IS a\n  multiple-line\r\n  com*ment*/", parser.GetTokenVal(l))
	parser.NextToken(l)
	r.Equal(Tokens.FROM, parser.GetToken(l))
	// plus
	l = parser.NewMySqlLexer("+")
	r.Equal(Tokens.PLUS, parser.NextToken(l))
	r.Equal("+", parser.GetTokenVal(l))
	// sub
	l = parser.NewMySqlLexer("-")
	r.Equal(Tokens.SUB, parser.NextToken(l))
	r.Equal("-", parser.GetTokenVal(l))
	// star
	l = parser.NewMySqlLexer("*")
	r.Equal(Tokens.STAR, parser.NextToken(l))
	r.Equal("*", parser.GetTokenVal(l))
	// slash
	l = parser.NewMySqlLexer("/")
	r.Equal(Tokens.SLASH, parser.NextToken(l))
	r.Equal("/", parser.GetTokenVal(l))
	// tilde
	l = parser.NewMySqlLexer("~")
	r.Equal(Tokens.TILDE, parser.NextToken(l))
	r.Equal("~", parser.GetTokenVal(l))
	// gt
	l = parser.NewMySqlLexer("> abc")
	r.Equal(Tokens.GT, parser.NextToken(l))
	r.Equal(">", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// gtGt
	l = parser.NewMySqlLexer(">> abc")
	r.Equal(Tokens.GT_GT, parser.NextToken(l))
	r.Equal(">>", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// gtEq
	l = parser.NewMySqlLexer(">= abc")
	r.Equal(Tokens.GT_EQ, parser.NextToken(l))
	r.Equal(">=", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// lt
	l = parser.NewMySqlLexer("< abc")
	r.Equal(Tokens.LT, parser.NextToken(l))
	r.Equal("<", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// ltGt
	l = parser.NewMySqlLexer("<> abc")
	r.Equal(Tokens.LT_GT, parser.NextToken(l))
	r.Equal("<>", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// ltLt
	l = parser.NewMySqlLexer("<< abc")
	r.Equal(Tokens.LT_LT, parser.NextToken(l))
	r.Equal("<<", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// ltEqGt
	l = parser.NewMySqlLexer("<=> abc")
	r.Equal(Tokens.LT_EQ_GT, parser.NextToken(l))
	r.Equal("<=>", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// ltEq
	l = parser.NewMySqlLexer("<= abc")
	r.Equal(Tokens.LT_EQ, parser.NextToken(l))
	r.Equal("<=", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// amp
	l = parser.NewMySqlLexer("& abc")
	r.Equal(Tokens.AMP, parser.NextToken(l))
	r.Equal("&", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// ampAmp
	l = parser.NewMySqlLexer("&& abc")
	r.Equal(Tokens.AMP_AMP, parser.NextToken(l))
	r.Equal("&&", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// at
	l = parser.NewMySqlLexer("@abc")
	r.Equal(Tokens.AT, parser.NextToken(l))
	r.Equal("@abc", parser.GetTokenVal(l))
	// atAt
	l = parser.NewMySqlLexer("@@abc")
	r.Equal(Tokens.AT_AT, parser.NextToken(l))
	r.Equal("@@abc", parser.GetTokenVal(l))
	// bar
	l = parser.NewMySqlLexer("| abc")
	r.Equal(Tokens.BAR, parser.NextToken(l))
	r.Equal("|", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// barBar
	l = parser.NewMySqlLexer("|| abc")
	r.Equal(Tokens.BAR_BAR, parser.NextToken(l))
	r.Equal("||", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// bang
	l = parser.NewMySqlLexer("! abc")
	r.Equal(Tokens.BANG, parser.NextToken(l))
	r.Equal("!", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// bangEq
	l = parser.NewMySqlLexer("!= abc")
	r.Equal(Tokens.BANG_EQ, parser.NextToken(l))
	r.Equal("!=", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// eq
	l = parser.NewMySqlLexer("= abc")
	r.Equal(Tokens.EQ, parser.NextToken(l))
	r.Equal("=", parser.GetTokenVal(l))
	r.Equal(Tokens.IDENTIFIER, parser.NextToken(l))
	r.Equal("abc", parser.GetTokenVal(l))
	// at
	l = parser.NewMySqlLexer("@abc")
	r.Equal(Tokens.AT, parser.NextToken(l))
	r.Equal("@abc", parser.GetTokenVal(l))
	l = parser.NewMySqlLexer("@'abc'")
	r.Equal(Tokens.AT, parser.NextToken(l))
	r.Equal("@'abc'", parser.GetTokenVal(l))
	l = parser.NewMySqlLexer("@@abc.fwef")
	r.Equal(Tokens.AT_AT, parser.NextToken(l))
	r.Equal("@@abc.fwef", parser.GetTokenVal(l))
	l = parser.NewMySqlLexer("@@'abc'")
	r.Equal(Tokens.AT_AT, parser.NextToken(l))
	r.Equal("@@'abc'", parser.GetTokenVal(l))
	// semi
	l = parser.NewMySqlLexer(";")
	r.Equal(Tokens.SEMI, parser.NextToken(l))
	r.Equal(";", parser.GetTokenVal(l))
	// colonEq
	l = parser.NewMySqlLexer(":=")
	r.Equal(Tokens.COLON_EQ, parser.NextToken(l))
	r.Equal(":=", parser.GetTokenVal(l))
	// percent
	l = parser.NewMySqlLexer("%")
	r.Equal(Tokens.PERCENT, parser.NextToken(l))
	r.Equal("%", parser.GetTokenVal(l))
	// caret
	l = parser.NewMySqlLexer("^")
	r.Equal(Tokens.CARET, parser.NextToken(l))
	r.Equal("^", parser.GetTokenVal(l))
	// ques
	l = parser.NewMySqlLexer("?")
	r.Equal(Tokens.QUES, parser.NextToken(l))
	r.Equal("?", parser.GetTokenVal(l))
	// jsonUnquote
	l = parser.NewMySqlLexer("->>")
	r.Equal(Tokens.JSON_UNQUOTE, parser.NextToken(l))
	r.Equal("->>", parser.GetTokenVal(l))
	// jsonExtract
	l = parser.NewMySqlLexer("->")
	r.Equal(Tokens.JSON_EXTRACT, parser.NextToken(l))
	r.Equal("->", parser.GetTokenVal(l))
	// test next token when already eoi
	l = parser.NewMySqlLexer("a")
	parser.NextToken(l)
	parser.NextToken(l)
	parser.NextToken(l)
	r.Equal(Tokens.EOI, parser.GetToken(l))
	// test eoi character in sql
	l = parser.NewMySqlLexer("a" + string(parser.Eoi) + "b")
	parser.NextToken(l)
	r.Equal(Tokens.IDENTIFIER, parser.GetToken(l))
	r.Equal("a"+string(parser.Eoi)+"b", parser.GetTokenVal(l))
}

func TestMySqlLexerPanic(t *testing.T) {
	validatePanic(t, ":+", `expected char '=', actual char '+'
:↪+↩`)
	validatePanic(t, "，", `illegal character '，'
↪，↩`)
	validatePanic(t, "`col1 `col2`", "need character ` to finish\n`col1↪↩ `col2`")
	validatePanic(t, "0b", "invalid binary number\n↪0b↩")
	validatePanic(t, "0x", "invalid hexadecimal number\n↪0x↩")
	validatePanic(t, "'abc", "need character ''' to finish\n'abc↪↩")
	validatePanic(t, "@'abc", "need character ' to finish\n@'abc↪↩")
	validatePanic(t, "/* comment", "need string '*/' to finish\n/* comment↪↩")
}
