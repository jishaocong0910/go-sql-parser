package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type ValueListSyntax struct {
	*Syntax__
	*ListSyntax__[ExprSyntax_]
}

func (this *ValueListSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitValueListSyntax(this)
}

func (this *ValueListSyntax) writeSql(builder *sqlBuilder) {
	this.Format = false
	this.ListSyntax__.writeSql(builder)
}

func NewValueListSyntax() *ValueListSyntax {
	s := &ValueListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[ExprSyntax_](s)
	s.ParenthesizeType = enum.ParenthesizeType_.TRUE
	return s
}
