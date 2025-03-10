package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type ValueListSyntax struct {
	*M_Syntax
	*M_ListSyntax[I_ExprSyntax]
}

func (this *ValueListSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitValueListSyntax(this)
}

func (this *ValueListSyntax) writeSql(builder *sqlBuilder) {
	this.Format = false
	this.M_ListSyntax.writeSql(builder)
}

func NewValueListSyntax() *ValueListSyntax {
	s := &ValueListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[I_ExprSyntax](s)
	s.ParenthesizeType = enum.ParenthesizeTypes.TRUE
	return s
}
