package ast

// 表引用
type I_TableReferenceSyntax interface {
	I_Syntax
	M_TableReferenceSyntax_() *M_TableReferenceSyntax
}

type M_TableReferenceSyntax struct {
	I I_TableReferenceSyntax
}

func (this *M_TableReferenceSyntax) M_TableReferenceSyntax_() *M_TableReferenceSyntax {
	return this
}

func ExtendTableReferenceSyntax(i I_TableReferenceSyntax) *M_TableReferenceSyntax {
	return &M_TableReferenceSyntax{I: i}
}
