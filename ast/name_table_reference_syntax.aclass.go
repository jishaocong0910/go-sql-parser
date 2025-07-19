package ast

// 表名称
type NameTableReferenceSyntax_ interface {
	NameTableReferenceSyntax_() *NameTableReferenceSyntax__
	TableReferenceSyntax_
}

type NameTableReferenceSyntax__ struct {
	I             NameTableReferenceSyntax_
	TableNameItem *TableNameItemSyntax
	Alias         IdentifierSyntax_
}

func (this *NameTableReferenceSyntax__) NameTableReferenceSyntax_() *NameTableReferenceSyntax__ {
	return this
}

func (this *NameTableReferenceSyntax__) accept(v_ Visitor_) {
	v_.visitor_().visitNameTableReferenceSyntax__(this)
}

func ExtendNameTableReferenceSyntax(i NameTableReferenceSyntax_) *NameTableReferenceSyntax__ {
	return &NameTableReferenceSyntax__{I: i}
}
