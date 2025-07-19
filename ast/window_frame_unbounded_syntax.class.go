package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type WindowFrameUnboundedSyntax struct {
	*Syntax__
	*WindowFrameExtentSyntax__
	*WindowFrameStartEndSyntax__
	Type enum.WindowFrameStartEndType
}

func (this *WindowFrameUnboundedSyntax) accept(Visitor_) {}

func (this *WindowFrameUnboundedSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("UNBOUNDED PRECEDING")
}

func NewWindowFrameUnboundedSyntax() *WindowFrameUnboundedSyntax {
	s := &WindowFrameUnboundedSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.WindowFrameExtentSyntax__ = ExtendWindowFrameExtentSyntax(s)
	s.WindowFrameStartEndSyntax__ = ExtendWindowFrameStartEndSyntax(s)
	return s
}
