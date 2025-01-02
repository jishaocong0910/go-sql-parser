package ast

// 标识符列表
type I_IdentifierListSyntax interface {
	I_ExprListSyntax[I_IdentifierSyntax]
	M_C7EE99CCA546() *M_IdentifierListSyntax
}

type M_IdentifierListSyntax struct {
	I I_ExprListSyntax[I_IdentifierSyntax]
}

func (this *M_IdentifierListSyntax) M_C7EE99CCA546() *M_IdentifierListSyntax {
	return this
}

func ExtendIdentifierListSyntax(i I_IdentifierListSyntax) *M_IdentifierListSyntax {
	return &M_IdentifierListSyntax{I: i}
}
