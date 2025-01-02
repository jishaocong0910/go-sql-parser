package ast

import "go-sql-parser/enum"

// 二元操作
type I_BinaryOperationSyntax interface {
	I_ExprSyntax
	M_71D4793003A9() *M_BinaryOperationSyntax
}

type M_BinaryOperationSyntax struct {
	I              I_BinaryOperationSyntax
	LeftOperand    I_ExprSyntax
	RightOperand   I_ExprSyntax
	BinaryOperator enum.BinaryOperator
}

func (this *M_BinaryOperationSyntax) M_71D4793003A9() *M_BinaryOperationSyntax {
	return this
}

func ExtendBinaryOperationSyntax(i I_BinaryOperationSyntax) *M_BinaryOperationSyntax {
	return &M_BinaryOperationSyntax{I: i}
}
