package ast

// 表名称
type I_NameTableReferenceSyntax interface {
	I_TableReferenceSyntax
	M_0E797D96D386() *M_NameTableReferenceSyntax
}

type M_NameTableReferenceSyntax struct {
	I             I_NameTableReferenceSyntax
	TableNameItem *TableNameItemSyntax
	Alias         I_IdentifierSyntax
}

func (this *M_NameTableReferenceSyntax) M_0E797D96D386() *M_NameTableReferenceSyntax {
	return this
}

func (this *M_NameTableReferenceSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitNameTableReferenceSyntax(this)
}

func ExtendNameTableReferenceSyntax(i I_NameTableReferenceSyntax) *M_NameTableReferenceSyntax {
	return &M_NameTableReferenceSyntax{I: i}
}
