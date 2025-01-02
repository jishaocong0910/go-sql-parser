package ast

import "go-sql-parser/enum"

type MySqlBinaryOperationSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_BinaryOperationSyntax
	ComparisonMode      enum.MySqlComparisonMode
	LikeEscape          *MySqlStringSyntax
	BetweenThirdOperand I_ExprSyntax
}

func (this *MySqlBinaryOperationSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlBinaryOperationSyntax(this)
}

func (this *MySqlBinaryOperationSyntax) writeSql(builder *sqlBuilder) {
	if this.Format {
		if enum.ParenthesizeTypes.Is(this.ParenthesizeType, enum.ParenthesizeTypes.TRUE) {
			this.Format = false
		} else if enum.OperatorTypes.Not(this.BinaryOperator.OperatorType, enum.OperatorTypes.LOGICAL) {
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_BinaryOperationSyntax = ExtendBinaryOperationSyntax(s)
	return s
}
