package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// 二元操作
type BinaryOperationSyntax_ interface {
	BinaryOperationSyntax_() *BinaryOperationSyntax__
	ExprSyntax_
}

type BinaryOperationSyntax__ struct {
	I                   BinaryOperationSyntax_
	LeftOperand         ExprSyntax_
	RightOperand        ExprSyntax_
	BinaryOperator      enum.BinaryOperator
	BetweenThirdOperand ExprSyntax_
}

func (this *BinaryOperationSyntax__) BinaryOperationSyntax_() *BinaryOperationSyntax__ {
	return this
}

func (this *BinaryOperationSyntax__) accept(v_ Visitor_) {
	v_.visitor_().visitBinaryOperationSyntax__(this)
}

func ExtendBinaryOperationSyntax(i BinaryOperationSyntax_) *BinaryOperationSyntax__ {
	return &BinaryOperationSyntax__{I: i}
}
