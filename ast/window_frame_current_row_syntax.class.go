package ast

type WindowFrameCurrentRowSyntax struct {
	*M_Syntax
	*M_WindowFrameExtentSyntax
	*M_WindowFrameStartEndSyntax
}

func (this *WindowFrameCurrentRowSyntax) accept(I_Visitor) {}

func (this *WindowFrameCurrentRowSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("CURRENT ROW")
}

func NewWindowFrameCurrentRowSyntax() *WindowFrameCurrentRowSyntax {
	s := &WindowFrameCurrentRowSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_WindowFrameExtentSyntax = ExtendWindowFrameExtentSyntax(s)
	s.M_WindowFrameStartEndSyntax = ExtendWindowFrameStartEndSyntax(s)
	return s
}
