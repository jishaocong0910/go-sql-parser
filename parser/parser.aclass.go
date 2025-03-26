package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jishaocong0910/go-sql-parser/ast"

	. "github.com/jishaocong0910/go-sql-parser/enum"

	"github.com/pkg/errors"
)

type i_Parser interface {
	m_Parser_() *m_Parser
	parseIStatementSyntax() ast.I_StatementSyntax
}

type m_Parser struct {
	i i_Parser
	// 词法器
	lexer i_Lexer
	// SQL中的参数占位符总数
	parameterCounts int
}

func (this *m_Parser) m_Parser_() *m_Parser {
	return this
}

func (this *m_Parser) nextToken() Token {
	return this.lexer.nextToken()
}

func (this *m_Parser) nextTokenIncludeComment() Token {
	return this.lexer.nextTokenIncludeComment()
}

func (this *m_Parser) prevToken() Token {
	return this.lexer.prevToken()
}

func (this *m_Parser) token() Token {
	return this.lexer.token()
}

func (this *m_Parser) tokenBeginPos() int {
	return this.lexer.tokenBeginPos()
}

func (this *m_Parser) tokenEndPos() int {
	return this.lexer.tokenEndPos()
}

func (this *m_Parser) prevTokenBeginPos() int {
	return this.lexer.prevTokenBeginPos()
}

func (this *m_Parser) prevTokenEndPos() int {
	return this.lexer.prevTokenEndPos()
}

func (this *m_Parser) tokenVal() string {
	return this.lexer.tokenVal()
}

func (this *m_Parser) tokenValUpper() string {
	return this.lexer.tokenValUpper()
}

func (this *m_Parser) reserved() bool {
	return this.lexer.reserved()
}

func (this *m_Parser) sql() string {
	return this.lexer.m_Lexer_().sql
}

func (this *m_Parser) saveCursor() *cursor {
	return this.lexer.saveCursor()
}

func (this *m_Parser) rollback(c *cursor) {
	this.lexer.rollback(c)
}

func (this *m_Parser) nextParameterIndex() int {
	this.parameterCounts++
	return this.parameterCounts
}

func (this *m_Parser) equalTokenVal(tokenVal string) bool {
	return strings.EqualFold(this.tokenVal(), tokenVal)
}

func (this *m_Parser) acceptAnyToken(tokens ...Token) {
	match := false
	if Tokens.Is(this.token(), tokens...) {
		match = true
	}
	if !match {
		var builder strings.Builder
		builder.WriteString("expected token: ")
		separator := ""
		for i := range tokens {
			builder.WriteString(separator)
			builder.WriteString(tokens[i].ID())
			separator = ", "
		}
		builder.WriteString(", actual token: ")
		builder.WriteString(this.token().ID())
		this.panicByToken(builder.String())
	}
}

func (this *m_Parser) acceptAnyTokenVal(tokenVals ...string) {
	match := false
	for i := range tokenVals {
		if strings.EqualFold(this.tokenVal(), tokenVals[i]) {
			match = true
			break
		}
	}
	if !match {
		var builder strings.Builder
		builder.WriteString("expected tokenVal: ")
		separator := ""
		for i := range tokenVals {
			builder.WriteString(separator)
			builder.WriteString(tokenVals[i])
			separator = ", "
		}
		builder.WriteString(", actual tokenVal: ")
		tokenVal := this.tokenVal()
		chars := []rune(tokenVal)
		if len(chars) == 1 {
			tokenVal = Character.CharDesc(chars[0])
		}
		builder.WriteString(tokenVal)
		this.panicByToken(builder.String())
	}
}

func (this *m_Parser) acceptExpectedOperandCount(expected int, operand ast.I_ExprSyntax) {
	// 小于0则表示无需比较，例如表达式为子查询，其查询列表为*，如SELECT * FROM tab1，由于*无法知道具体列数，遂跳过
	actual := operand.OperandCount()
	if !(expected > 0 && actual > 0) {
		return
	}
	if expected != actual {
		this.panicBySyntax(operand, "operand should contain "+strconv.Itoa(expected)+" Column(s), but found "+strconv.Itoa(actual))
	}
}

func (this *m_Parser) acceptEqualOperandCount(leftOperand, rightOperand ast.I_ExprSyntax, inOperator bool) {
	if leftOperand.IsExprList() {
		if rightOperand.IsExprList() {
			if inOperator {
				for i := 0; i < rightOperand.ExprLen(); i++ {
					this.acceptEqualOperandCount(leftOperand, rightOperand.GetExpr(i), false)
				}
			} else {
				this.acceptExpectedOperandCount(leftOperand.OperandCount(), rightOperand)
				for i := 0; i < leftOperand.ExprLen(); i++ {
					this.acceptEqualOperandCount(leftOperand.GetExpr(i), rightOperand.GetExpr(i), false)
				}
			}
		} else {
			this.acceptExpectedOperandCount(leftOperand.OperandCount(), rightOperand)
		}
	} else if rightOperand.IsExprList() {
		for i := 0; i < rightOperand.ExprLen(); i++ {
			this.acceptExpectedOperandCount(1, rightOperand.GetExpr(i))
		}
	} else {
		this.acceptExpectedOperandCount(leftOperand.OperandCount(), rightOperand)
	}
}

func (this *m_Parser) parenthesizingSyntax(is ast.I_Syntax) {
	s := is.M_Syntax_()
	if ParenthesizeTypes.Is(s.ParenthesizeType, ParenthesizeTypes.NOT_SUPPORT) {
		this.panicBySyntax(is, "this syntax cannot be parenthesized")
	}
	s.ParenthesizeType = ParenthesizeTypes.TRUE
}

func (this *m_Parser) setBeginPos(is ast.I_Syntax, pos int) {
	is.M_Syntax_().BeginPos = pos
}
func (this *m_Parser) setBeginPosDefault(is ast.I_Syntax) {
	is.M_Syntax_().BeginPos = this.tokenBeginPos()
}

func (this *m_Parser) setEndPos(is ast.I_Syntax, pos int) {
	is.M_Syntax_().EndPos = pos
}

func (this *m_Parser) setEndPosDefault(is ast.I_Syntax) {
	is.M_Syntax_().EndPos = this.prevTokenEndPos()
}

func (this *m_Parser) panicByUnexpectedToken() {
	this.panicByToken("unexpected token")
}

func (this *m_Parser) panicBySyntax(is ast.I_Syntax, msg string, a ...any) {
	s := is.M_Syntax_()
	this.panic(s.BeginPos, s.EndPos, msg, a...)
}

func (this *m_Parser) panicByToken(msg string, a ...any) {
	this.panic(this.tokenBeginPos(), this.tokenEndPos(), msg, a...)
}

func (this *m_Parser) panic(beginPos, endPos int, msg string, a ...any) {
	var builder strings.Builder
	if msg != "" {
		msg = fmt.Sprintf(msg, a...)
		builder.WriteString(msg)
	}
	builder.WriteString("\n")

	chars := this.lexer.m_Lexer_().chars
	if beginPos < len(chars) {
		for i := range chars {
			c := chars[i]
			if i == beginPos {
				builder.WriteString("↪")
			}
			builder.WriteRune(c)
			if i == endPos-1 {
				builder.WriteString("↩")
			}
		}
	} else {
		builder.WriteString(this.sql())
		builder.WriteString("↪↩")
	}
	panic(parseError(builder.String()))
}

func extendParser(i i_Parser, lexer i_Lexer) *m_Parser {
	p := &m_Parser{i: i, lexer: lexer}
	p.lexer.nextToken()
	return p
}

type parseError string

func (e parseError) Error() string {
	return string(e)
}

func Parse(d Dialect, sql string) (is ast.I_StatementSyntax, err error) {
	var ip i_Parser
	switch d.ID() {
	case Dialects.MYSQL.ID():
		ip = newMySqlParser(sql)
	default:
		return nil, fmt.Errorf("not supported database type for '%s'", d.Name)
	}
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(parseError); ok {
				err = e
			} else {
				err = errors.WithStack(fmt.Errorf("%v", r))
			}
		}
	}()
	is = ip.parseIStatementSyntax()
	return
}
