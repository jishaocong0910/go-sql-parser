package ast

// 具有WHERE子句的语法
type I_HaveWhereSyntax interface {
	I_Syntax
	M_HaveWhereSyntax_() *M_HaveWhereSyntax
}

type M_HaveWhereSyntax struct {
	I     I_HaveWhereSyntax
	Where *WhereSyntax
}

func (this *M_HaveWhereSyntax) M_HaveWhereSyntax_() *M_HaveWhereSyntax {
	return this
}

func ExtendHaveWhereSyntax(i I_HaveWhereSyntax) *M_HaveWhereSyntax {
	return &M_HaveWhereSyntax{I: i}
}
