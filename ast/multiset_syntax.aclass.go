package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"
)

// 多结果集语法，例如UNION、EXCEPT、INTERSECT
type MultisetSyntax_ interface {
	MultisetSyntax_() *MultisetSyntax__
	QuerySyntax_
}

type MultisetSyntax__ struct {
	I                MultisetSyntax_
	LeftQuery        QuerySyntax_
	RightQuery       QuerySyntax_
	MultisetOperator enum.MultisetOperator
	AggregateOption  enum.AggregateOption
	OrderBy          *OrderBySyntax
}

func (this *MultisetSyntax__) MultisetSyntax_() *MultisetSyntax__ {
	return this
}

func (this *MultisetSyntax__) accept(v_ Visitor_) {
	v_.visitor_().visitMultisetSyntax__(this)
}

func (this *MultisetSyntax__) OperandCount() int {
	return this.LeftQuery.OperandCount()
}

func ExtendMultisetSyntax(i MultisetSyntax_) *MultisetSyntax__ {
	return &MultisetSyntax__{I: i}
}
