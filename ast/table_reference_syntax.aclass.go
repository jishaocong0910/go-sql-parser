package ast

// 表引用
type I_TableReferenceSyntax interface {
	I_Syntax
	M_4C601394116B() *M_TableReferenceSyntax
}

type M_TableReferenceSyntax struct {
	I I_TableReferenceSyntax
}

func (this *M_TableReferenceSyntax) M_4C601394116B() *M_TableReferenceSyntax {
	return this
}

func ExtendTableReferenceSyntax(i I_TableReferenceSyntax) *M_TableReferenceSyntax {
	return &M_TableReferenceSyntax{I: i}
}
