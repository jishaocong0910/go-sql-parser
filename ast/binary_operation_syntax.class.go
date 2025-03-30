package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type BinaryOperationSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_BinaryOperationSyntax
}

func (this *BinaryOperationSyntax) writeSql(builder *sqlBuilder) {
	if this.Format {
		if enum.ParenthesizeTypes.Is(this.ParenthesizeType, enum.ParenthesizeTypes.TRUE) {
			this.Format = false
		} else if enum.OperatorTypes.Not(this.BinaryOperator.OperatorType, enum.OperatorTypes.LOGICAL) {
			this.Format = false
		}
	}
	builder.writeSyntax(this.LeftOperand)
	builder.writeSpaceOrLfIndent(this, this.BinaryOperator.Symbol, " ")
	builder.writeSyntax(this.RightOperand)
	if this.BetweenThirdOperand != nil {
		builder.writeStr(" AND ")
		builder.writeSyntax(this.BetweenThirdOperand)
	}
}

func NewBinaryOperationSyntax() *BinaryOperationSyntax {
	s := &BinaryOperationSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_BinaryOperationSyntax = ExtendBinaryOperationSyntax(s)
	return s
}
