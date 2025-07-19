package ast

// 别名
type AliasSyntax_ interface {
	AliasSyntax_() *AliasSyntax__
	Syntax_

	AliasName() string
}

type AliasSyntax__ struct {
	I AliasSyntax_
}

func (this *AliasSyntax__) AliasSyntax_() *AliasSyntax__ {
	return this
}

func ExtendAliasSyntax(i AliasSyntax_) *AliasSyntax__ {
	return &AliasSyntax__{I: i}
}
