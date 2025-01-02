package ast

type OrderingItemListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*OrderingItemSyntax]
}

func NewOrderingItemListSyntax() *OrderingItemListSyntax {
	s := &OrderingItemListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*OrderingItemSyntax](s)
	return s
}
