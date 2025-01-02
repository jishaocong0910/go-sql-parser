package ast

type CaseWhenItemListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*CaseWhenItemSyntax]
}

func NewCaseWhenItemListSyntax() *CaseWhenItemListSyntax {
	s := &CaseWhenItemListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*CaseWhenItemSyntax](s)
	s.separator = ""
	return s
}
