package ast

// 别名
type I_AliasSyntax interface {
	I_Syntax
	M_AliasSyntax_() *M_AliasSyntax
	AliasName() string
}

type M_AliasSyntax struct {
	I I_AliasSyntax
}

func (this *M_AliasSyntax) M_AliasSyntax_() *M_AliasSyntax {
	return this
}

func ExtendAliasSyntax(i I_AliasSyntax) *M_AliasSyntax {
	return &M_AliasSyntax{I: i}
}
