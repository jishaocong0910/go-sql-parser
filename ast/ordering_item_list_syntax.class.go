package ast

type OrderingItemListSyntax struct {
	*Syntax__
	*ListSyntax__[*OrderingItemSyntax]
}

func NewOrderingItemListSyntax() *OrderingItemListSyntax {
	s := &OrderingItemListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*OrderingItemSyntax](s)
	return s
}
