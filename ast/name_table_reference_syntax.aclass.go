package ast

// 表名称
type I_NameTableReferenceSyntax interface {
	I_TableReferenceSyntax
	M_NameTableReferenceSyntax_() *M_NameTableReferenceSyntax
}

type M_NameTableReferenceSyntax struct {
	I             I_NameTableReferenceSyntax
	TableNameItem *TableNameItemSyntax
	Alias         I_IdentifierSyntax
}

func (this *M_NameTableReferenceSyntax) M_NameTableReferenceSyntax_() *M_NameTableReferenceSyntax {
	return this
}

func (this *M_NameTableReferenceSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitNameTableReferenceSyntax(this)
}

func ExtendNameTableReferenceSyntax(i I_NameTableReferenceSyntax) *M_NameTableReferenceSyntax {
	return &M_NameTableReferenceSyntax{I: i}
}
