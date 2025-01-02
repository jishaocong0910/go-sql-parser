package ast

type SelectItemListSyntax struct {
	*M_Syntax
	*M_ListSyntax[I_SelectItemSyntax]
	HasAllColumn bool
}

func (this *SelectItemListSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitSelectItemListSyntax(this)
}

func (this *SelectItemListSyntax) writeSql(builder *sqlBuilder) {
	this.M_ListSyntax.writeSql(builder)
}

func NewSelectItemListSyntax() *SelectItemListSyntax {
	s := &SelectItemListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[I_SelectItemSyntax](s)
	return s
}
