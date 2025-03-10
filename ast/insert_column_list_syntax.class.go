package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type InsertColumnListSyntax struct {
	*M_Syntax
	*M_ListSyntax[I_IdentifierSyntax]
}

func (this *InsertColumnListSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitInsertColumnListSyntax(this)
}

func NewInsertColumnListSyntax() *InsertColumnListSyntax {
	s := &InsertColumnListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[I_IdentifierSyntax](s)
	s.ParenthesizeType = enum.ParenthesizeTypes.TRUE
	return s
}
