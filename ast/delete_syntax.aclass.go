package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// DELETE语法
type DeleteSyntax_ interface {
	DeleteSyntax_() *DeleteSyntax__
	StatementSyntax_
	HaveWhereSyntax_
}

type DeleteSyntax__ struct {
	I              DeleteSyntax_
	TableReference TableReferenceSyntax_
	Hint           *HintSyntax
}

func (this *DeleteSyntax__) DeleteSyntax_() *DeleteSyntax__ {
	return this
}

func ExtendDeleteSyntax(i DeleteSyntax_) *DeleteSyntax__ {
	i.Syntax_().ParenthesizeType = enum.ParenthesizeType_.NOT_SUPPORT
	return &DeleteSyntax__{I: i}
}
