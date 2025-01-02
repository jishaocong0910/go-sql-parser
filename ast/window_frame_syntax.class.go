package ast

import "go-sql-parser/enum"

type WindowFrameSyntax struct {
	*M_Syntax
	Unit   enum.WindowFrameUnit
	Extent I_WindowFrameExtentSyntax
}

func (this *WindowFrameSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitWindowFrameSyntax(this)
}

func (this *WindowFrameSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Unit.Sql)
	builder.writeSpace()
	builder.writeSyntax(this.Extent)
}

func NewWindowFrameSyntax() *WindowFrameSyntax {
	s := &WindowFrameSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
