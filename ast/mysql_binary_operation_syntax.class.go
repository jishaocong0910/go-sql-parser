package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlBinaryOperationSyntax struct {
	*Syntax__
	*ExprSyntax__
	*BinaryOperationSyntax__
	ComparisonMode enum.MySqlComparisonMode
	LikeEscape     *MySqlStringSyntax
}

func (this *MySqlBinaryOperationSyntax) writeSql(builder *sqlBuilder) {
	if this.Format {
		if enum.ParenthesizeType_.Is(this.ParenthesizeType, enum.ParenthesizeType_.TRUE) {
			this.Format = false
		} else if enum.OperatorType_.Not(this.BinaryOperator.OperatorType, enum.OperatorType_.LOGICAL) {
			this.Format = false
		}
	}
	builder.writeSyntax(this.LeftOperand)
	builder.writeSpaceOrLfIndent(this, this.BinaryOperator.Symbol, " ")
	if !this.ComparisonMode.Undefined() {
		builder.writeStr(this.ComparisonMode.Sql)
		builder.writeSpace()
	}
	builder.writeSyntax(this.RightOperand)
	if this.LikeEscape != nil {
		builder.writeStr(" ESCAPE ")
		builder.writeSyntax(this.LikeEscape)
	}
	if this.BetweenThirdOperand != nil {
		builder.writeStr(" AND ")
		builder.writeSyntax(this.BetweenThirdOperand)
	}
}

func NewMySqlBinaryOperationSyntax() *MySqlBinaryOperationSyntax {
	s := &MySqlBinaryOperationSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.BinaryOperationSyntax__ = ExtendBinaryOperationSyntax(s)
	return s
}
