package ast

type WindowFrameCurrentRowSyntax struct {
	*Syntax__
	*WindowFrameExtentSyntax__
	*WindowFrameStartEndSyntax__
}

func (this *WindowFrameCurrentRowSyntax) accept(Visitor_) {}

func (this *WindowFrameCurrentRowSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("CURRENT ROW")
}

func NewWindowFrameCurrentRowSyntax() *WindowFrameCurrentRowSyntax {
	s := &WindowFrameCurrentRowSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.WindowFrameExtentSyntax__ = ExtendWindowFrameExtentSyntax(s)
	s.WindowFrameStartEndSyntax__ = ExtendWindowFrameStartEndSyntax(s)
	return s
}
