package ast

type MySqlIndexHintListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*MySqlIndexHintSyntax]
}

func (this *MySqlIndexHintListSyntax) accept(I_Visitor) {}

func NewMySqlIndexHintListSyntax() *MySqlIndexHintListSyntax {
	s := &MySqlIndexHintListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*MySqlIndexHintSyntax](s)
	s.separator = ""
	return s
}
