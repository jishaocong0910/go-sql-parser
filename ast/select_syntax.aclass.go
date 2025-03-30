package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// SELECT语法
type I_SelectSyntax interface {
	I_QuerySyntax
	I_HaveWhereSyntax
	M_SelectSyntax_() *M_SelectSyntax
}

type M_SelectSyntax struct {
	I               I_SelectSyntax
	AggregateOption enum.AggregateOption
	SelectItemList  *SelectItemListSyntax
	TableReference  I_TableReferenceSyntax
	GroupBy         I_GroupBySyntax
	Having          *HavingSyntax
	NamedWindowList *NamedWindowsListSyntax
	OrderBy         *OrderBySyntax
	Hint            *HintSyntax
}

func (this *M_SelectSyntax) M_SelectSyntax_() *M_SelectSyntax {
	return this
}

func (this *M_SelectSyntax) OperandCount() int {
	if this.SelectItemList.HasAllColumn {
		return 0
	}
	return this.SelectItemList.Len()
}

func ExtendSelectSyntax(i I_SelectSyntax) *M_SelectSyntax {
	return &M_SelectSyntax{I: i}
}
