package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type WindowFrameSyntax struct {
	*Syntax__
	Unit   enum.WindowFrameUnit
	Extent WindowFrameExtentSyntax_
}

func (this *WindowFrameSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitWindowFrameSyntax(this)
}

func (this *WindowFrameSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Unit.Sql)
	builder.writeSpace()
	builder.writeSyntax(this.Extent)
}

func NewWindowFrameSyntax() *WindowFrameSyntax {
	s := &WindowFrameSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
