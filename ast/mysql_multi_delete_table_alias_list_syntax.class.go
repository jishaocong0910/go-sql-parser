package ast

type MySqlMultiDeleteTableAliasListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*MySqlMultiDeleteTableAliasSyntax]
}

func NewMySqlMultiDeleteTableAliasListSyntax() *MySqlMultiDeleteTableAliasListSyntax {
	s := &MySqlMultiDeleteTableAliasListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*MySqlMultiDeleteTableAliasSyntax](s)
	return s
}
