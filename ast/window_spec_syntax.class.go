package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"

	o "github.com/jishaocong0910/go-object-util"
)

type WindowSpecSyntax struct {
	*Syntax__
	*OverWindowSyntax__
	Name        IdentifierSyntax_
	PartitionBy *PartitionBySyntax // 标准SQL要求 PARTITION BY 语法后接字段名，MySQL扩展至支持表达式
	OrderBy     *OrderBySyntax
	Frame       *WindowFrameSyntax
}

func (this *WindowSpecSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitWindowSpecSyntax(this)
}

func (this *WindowSpecSyntax) writeSql(builder *sqlBuilder) {
	first := true
	w := func(s_ Syntax_, format bool) {
		if o.IsNull(s_) {
			return
		}
		if !first {
			builder.writeSpace()
		}
		builder.writeSyntaxWithFormat(s_, format)
		first = false
	}
	w(this.Name, true)
	w(this.PartitionBy, true)
	w(this.OrderBy, false)
	w(this.Frame, true)
}

func NewWindowSpecSyntax() *WindowSpecSyntax {
	s := &WindowSpecSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.OverWindowSyntax__ = ExtendOverWindowSyntax(s)
	s.ParenthesizeType = enum.ParenthesizeType_.TRUE
	return s
}
