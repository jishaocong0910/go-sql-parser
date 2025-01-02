package ast

// 别名
type I_AliasSyntax interface {
	I_Syntax
	M_7869262892CE() *M_AliasSyntax
	AliasName() string
}

type M_AliasSyntax struct {
	I I_AliasSyntax
}

func (this *M_AliasSyntax) M_7869262892CE() *M_AliasSyntax {
	return this
}

func ExtendAliasSyntax(i I_AliasSyntax) *M_AliasSyntax {
	return &M_AliasSyntax{I: i}
}
