package ast

type SelectItemListSyntax struct {
	*Syntax__
	*ListSyntax__[SelectItemSyntax_]
	HasAllColumn bool
}

func (this *SelectItemListSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitSelectItemListSyntax(this)
}

func (this *SelectItemListSyntax) writeSql(builder *sqlBuilder) {
	this.ListSyntax__.writeSql(builder)
}

func NewSelectItemListSyntax() *SelectItemListSyntax {
	s := &SelectItemListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[SelectItemSyntax_](s)
	return s
}
