package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type InsertColumnListSyntax struct {
	*Syntax__
	*ListSyntax__[IdentifierSyntax_]
}

func (this *InsertColumnListSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitInsertColumnListSyntax(this)
}

func NewInsertColumnListSyntax() *InsertColumnListSyntax {
	s := &InsertColumnListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[IdentifierSyntax_](s)
	s.ParenthesizeType = enum.ParenthesizeType_.TRUE
	return s
}
