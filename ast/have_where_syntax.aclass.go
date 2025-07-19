package ast

// 具有WHERE子句的语法
type HaveWhereSyntax_ interface {
	HaveWhereSyntax_() *HaveWhereSyntax__
	Syntax_
}

type HaveWhereSyntax__ struct {
	I     HaveWhereSyntax_
	Where *WhereSyntax
}

func (this *HaveWhereSyntax__) HaveWhereSyntax_() *HaveWhereSyntax__ {
	return this
}

func ExtendHaveWhereSyntax(i HaveWhereSyntax_) *HaveWhereSyntax__ {
	return &HaveWhereSyntax__{I: i}
}
