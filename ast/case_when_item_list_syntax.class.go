package ast

type CaseWhenItemListSyntax struct {
	*Syntax__
	*ListSyntax__[*CaseWhenItemSyntax]
}

func NewCaseWhenItemListSyntax() *CaseWhenItemListSyntax {
	s := &CaseWhenItemListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*CaseWhenItemSyntax](s)
	s.separator = ""
	return s
}
