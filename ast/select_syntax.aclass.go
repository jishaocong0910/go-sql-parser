package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// SELECT语法
type SelectSyntax_ interface {
	SelectSyntax_() *SelectSyntax__
	QuerySyntax_
	HaveWhereSyntax_
}

type SelectSyntax__ struct {
	I               SelectSyntax_
	AggregateOption enum.AggregateOption
	SelectItemList  *SelectItemListSyntax
	TableReference  TableReferenceSyntax_
	GroupBy         GroupBySyntax_
	Having          *HavingSyntax
	NamedWindowList *NamedWindowsListSyntax
	OrderBy         *OrderBySyntax
	Hint            *HintSyntax
}

func (this *SelectSyntax__) SelectSyntax_() *SelectSyntax__ {
	return this
}

func (this *SelectSyntax__) OperandCount() int {
	if this.SelectItemList.HasAllColumn {
		return 0
	}
	return this.SelectItemList.Len()
}

func ExtendSelectSyntax(i SelectSyntax_) *SelectSyntax__ {
	return &SelectSyntax__{I: i}
}
