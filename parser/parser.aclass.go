package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jishaocong0910/go-sql-parser/ast"

	. "github.com/jishaocong0910/go-sql-parser/enum"

	"github.com/pkg/errors"
)

type Parser_ interface {
	parser() *parser__
	parseStatementSyntax_() ast.StatementSyntax_
}

type parser__ struct {
	i Parser_
	// 词法器
	lexer lexer_
	// SQL中的参数占位符总数
	parameterCounts int
}

func (this *parser__) parser() *parser__ {
	return this
}

func (this *parser__) nextToken() Token {
	return this.lexer.nextToken()
}

func (this *parser__) nextTokenIncludeComment() Token {
	return this.lexer.nextTokenIncludeComment()
}

func (this *parser__) prevToken() Token {
	return this.lexer.prevToken()
}

func (this *parser__) token() Token {
	return this.lexer.token()
}

func (this *parser__) tokenBeginPos() int {
	return this.lexer.tokenBeginPos()
}

func (this *parser__) tokenEndPos() int {
	return this.lexer.tokenEndPos()
}

func (this *parser__) prevTokenBeginPos() int {
	return this.lexer.prevTokenBeginPos()
}

func (this *parser__) prevTokenEndPos() int {
	return this.lexer.prevTokenEndPos()
}

func (this *parser__) tokenVal() string {
	return this.lexer.tokenVal()
}

func (this *parser__) tokenValUpper() string {
	return this.lexer.tokenValUpper()
}

func (this *parser__) reserved() bool {
	return this.lexer.reserved()
}

func (this *parser__) sql() string {
	return this.lexer.lexer_().sql
}

func (this *parser__) saveCursor() *cursor {
	return this.lexer.saveCursor()
}

func (this *parser__) rollback(c *cursor) {
	this.lexer.rollback(c)
}

func (this *parser__) nextParameterIndex() int {
	this.parameterCounts++
	return this.parameterCounts
}

func (this *parser__) equalTokenVal(tokenVal string) bool {
	return strings.EqualFold(this.tokenVal(), tokenVal)
}

func (this *parser__) acceptAnyToken(tokens ...Token) {
	match := false
	if Token_.Is(this.token(), tokens...) {
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

func (this *parser__) acceptAnyTokenVal(tokenVals ...string) {
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

func (this *parser__) acceptExpectedOperandCount(expected int, operand ast.ExprSyntax_) {
	// 小于0则表示无需比较，例如表达式为子查询，其查询列表为*，如SELECT * FROM tab1，由于*无法知道具体列数，遂跳过
	actual := operand.OperandCount()
	if !(expected > 0 && actual > 0) {
		return
	}
	if expected != actual {
		this.panicBySyntax(operand, "operand should contain "+strconv.Itoa(expected)+" Column(s), but found "+strconv.Itoa(actual))
	}
}

func (this *parser__) acceptEqualOperandCount(leftOperand, rightOperand ast.ExprSyntax_, inOperator bool) {
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

func (this *parser__) parenthesizingSyntax(is ast.Syntax_) {
	s := is.Syntax_()
	if ParenthesizeType_.Is(s.ParenthesizeType, ParenthesizeType_.NOT_SUPPORT) {
		this.panicBySyntax(is, "this syntax cannot be parenthesized")
	}
	s.ParenthesizeType = ParenthesizeType_.TRUE
}

func (this *parser__) setBeginPos(is ast.Syntax_, pos int) {
	is.Syntax_().BeginPos = pos
}
func (this *parser__) setBeginPosDefault(is ast.Syntax_) {
	is.Syntax_().BeginPos = this.tokenBeginPos()
}

func (this *parser__) setEndPos(is ast.Syntax_, pos int) {
	is.Syntax_().EndPos = pos
}

func (this *parser__) setEndPosDefault(is ast.Syntax_) {
	is.Syntax_().EndPos = this.prevTokenEndPos()
}

func (this *parser__) panicByUnexpectedToken() {
	this.panicByToken("unexpected token")
}

func (this *parser__) panicBySyntax(is ast.Syntax_, msg string, a ...any) {
	s := is.Syntax_()
	this.panic(s.BeginPos, s.EndPos, msg, a...)
}

func (this *parser__) panicByToken(msg string, a ...any) {
	this.panic(this.tokenBeginPos(), this.tokenEndPos(), msg, a...)
}

func (this *parser__) panic(beginPos, endPos int, msg string, a ...any) {
	var builder strings.Builder
	if msg != "" {
		msg = fmt.Sprintf(msg, a...)
		builder.WriteString(msg)
	}
	builder.WriteString("\n")

	chars := this.lexer.lexer_().chars
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

func extendParser(i Parser_, lexer lexer_) *parser__ {
	p := &parser__{i: i, lexer: lexer}
	p.lexer.nextToken()
	return p
}

type parseError string

func (e parseError) Error() string {
	return string(e)
}

func Parse(d Dialect, sql string) (s_ ast.StatementSyntax_, err error) {
	var p_ Parser_
	switch d.ID() {
	case Dialect_.MYSQL.ID():
		p_ = newMySqlParser(sql)
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
	s_ = p_.parseStatementSyntax_()
	return
}
