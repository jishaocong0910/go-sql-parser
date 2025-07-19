package ast

type NamedWindowsListSyntax struct {
	*Syntax__
	*ListSyntax__[*NamedWindowsSyntax]
}

func (this *NamedWindowsListSyntax) writeSql(builder *sqlBuilder) {
	this.Format = false
	this.ListSyntax__.writeSql(builder)
}

func NewNamedWindowsListSyntax() *NamedWindowsListSyntax {
	s := &NamedWindowsListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*NamedWindowsSyntax](s)
	return s
}
