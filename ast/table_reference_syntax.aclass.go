package ast

// 表引用
type TableReferenceSyntax_ interface {
	TableReferenceSyntax_() *TableReferenceSyntax__
	Syntax_
}

type TableReferenceSyntax__ struct {
	I TableReferenceSyntax_
}

func (this *TableReferenceSyntax__) TableReferenceSyntax_() *TableReferenceSyntax__ {
	return this
}

func ExtendTableReferenceSyntax(i TableReferenceSyntax_) *TableReferenceSyntax__ {
	return &TableReferenceSyntax__{I: i}
}
