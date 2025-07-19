package ast

type MySqlMultiDeleteTableAliasListSyntax struct {
	*Syntax__
	*ListSyntax__[*MySqlMultiDeleteTableAliasSyntax]
}

func NewMySqlMultiDeleteTableAliasListSyntax() *MySqlMultiDeleteTableAliasListSyntax {
	s := &MySqlMultiDeleteTableAliasListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*MySqlMultiDeleteTableAliasSyntax](s)
	return s
}
