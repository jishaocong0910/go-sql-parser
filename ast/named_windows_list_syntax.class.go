package ast

type NamedWindowsListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*NamedWindowsSyntax]
}

func (this *NamedWindowsListSyntax) writeSql(builder *sqlBuilder) {
	this.Format = false
	this.M_ListSyntax.writeSql(builder)
}

func NewNamedWindowsListSyntax() *NamedWindowsListSyntax {
	s := &NamedWindowsListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*NamedWindowsSyntax](s)
	return s
}
