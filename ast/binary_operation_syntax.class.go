package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type BinaryOperationSyntax struct {
	*Syntax__
	*ExprSyntax__
	*BinaryOperationSyntax__
}

func (this *BinaryOperationSyntax) writeSql(builder *sqlBuilder) {
	if this.Format {
		if enum.ParenthesizeType_.Is(this.ParenthesizeType, enum.ParenthesizeType_.TRUE) {
			this.Format = false
		} else if enum.OperatorType_.Not(this.BinaryOperator.OperatorType, enum.OperatorType_.LOGICAL) {
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
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.BinaryOperationSyntax__ = ExtendBinaryOperationSyntax(s)
	return s
}
