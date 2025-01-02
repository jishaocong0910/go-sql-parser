package ast

import (
	"go-sql-parser/enum"

	o "github.com/jishaocong0910/go-object"
)

type WindowSpecSyntax struct {
	*M_Syntax
	*M_OverWindowSyntax
	Name        I_IdentifierSyntax
	PartitionBy *PartitionBySyntax // 标准SQL要求 PARTITION BY 语法后接字段名，MySQL扩展至支持表达式
	OrderBy     *OrderBySyntax
	Frame       *WindowFrameSyntax
}

func (this *WindowSpecSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitWindowSpecSyntax(this)
}

func (this *WindowSpecSyntax) writeSql(builder *sqlBuilder) {
	first := true
	w := func(is I_Syntax, format bool) {
		if o.IsNull(is) {
			return
		}
		if !first {
			builder.writeSpace()
		}
		builder.writeSyntaxWithFormat(is, format)
		first = false
	}
	w(this.Name, true)
	w(this.PartitionBy, true)
	w(this.OrderBy, false)
	w(this.Frame, true)
}

func NewWindowSpecSyntax() *WindowSpecSyntax {
	s := &WindowSpecSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_OverWindowSyntax = ExtendOverWindowSyntax(s)
	s.ParenthesizeType = enum.ParenthesizeTypes.TRUE
	return s
}
