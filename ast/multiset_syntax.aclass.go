package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"
)

// 多结果集语法，例如UNION、EXCEPT、INTERSECT
type I_MultisetSyntax interface {
	I_QuerySyntax
	M_MultisetSyntax_() *M_MultisetSyntax
}

type M_MultisetSyntax struct {
	I                I_MultisetSyntax
	LeftQuery        I_QuerySyntax
	RightQuery       I_QuerySyntax
	MultisetOperator enum.MultisetOperator
	AggregateOption  enum.AggregateOption
	OrderBy          *OrderBySyntax
}

func (this *M_MultisetSyntax) M_MultisetSyntax_() *M_MultisetSyntax {
	return this
}

func (this *M_MultisetSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitMultisetSyntax(this)
}

func (this *M_MultisetSyntax) OperandCount() int {
	return this.LeftQuery.OperandCount()
}

func ExtendMultisetSyntax(i I_MultisetSyntax) *M_MultisetSyntax {
	return &M_MultisetSyntax{I: i}
}
