package ast

import "go-sql-parser/enum"

type WindowFrameUnboundedSyntax struct {
	*M_Syntax
	*M_WindowFrameExtentSyntax
	*M_WindowFrameStartEndSyntax
	Type enum.WindowFrameStartEndType
}

func (this *WindowFrameUnboundedSyntax) accept(I_Visitor) {}

func (this *WindowFrameUnboundedSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("UNBOUNDED PRECEDING")
}

func NewWindowFrameUnboundedSyntax() *WindowFrameUnboundedSyntax {
	s := &WindowFrameUnboundedSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_WindowFrameExtentSyntax = ExtendWindowFrameExtentSyntax(s)
	s.M_WindowFrameStartEndSyntax = ExtendWindowFrameStartEndSyntax(s)
	return s
}
