package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// 二元操作
type I_BinaryOperationSyntax interface {
	I_ExprSyntax
	M_BinaryOperationSyntax_() *M_BinaryOperationSyntax
}

type M_BinaryOperationSyntax struct {
	I                   I_BinaryOperationSyntax
	LeftOperand         I_ExprSyntax
	RightOperand        I_ExprSyntax
	BinaryOperator      enum.BinaryOperator
	BetweenThirdOperand I_ExprSyntax
}

func (this *M_BinaryOperationSyntax) M_BinaryOperationSyntax_() *M_BinaryOperationSyntax {
	return this
}

func (this *M_BinaryOperationSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitBinaryOperationSyntax(this)
}

func ExtendBinaryOperationSyntax(i I_BinaryOperationSyntax) *M_BinaryOperationSyntax {
	return &M_BinaryOperationSyntax{I: i}
}
