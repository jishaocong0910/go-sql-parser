package ast

// 标识符列表
type IdentifierListSyntax_ interface {
	IdentifierListSyntax_() *IdentifierListSyntax__
	ExprListSyntax_[IdentifierSyntax_]
}

type IdentifierListSyntax__ struct {
	I ExprListSyntax_[IdentifierSyntax_]
}

func (this *IdentifierListSyntax__) IdentifierListSyntax_() *IdentifierListSyntax__ {
	return this
}

func ExtendIdentifierListSyntax(i IdentifierListSyntax_) *IdentifierListSyntax__ {
	return &IdentifierListSyntax__{I: i}
}
