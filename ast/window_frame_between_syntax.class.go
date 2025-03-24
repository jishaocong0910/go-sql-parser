package ast

type WindowFrameBetweenSyntax struct {
	*M_Syntax
	*M_WindowFrameExtentSyntax
	Start I_WindowFrameStartEndSyntax
	End   I_WindowFrameStartEndSyntax
}

func (this *WindowFrameBetweenSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitWindowFrameBetweenSyntax(this)
}

func (this *WindowFrameBetweenSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("BETWEEN ")
	builder.writeSyntax(this.Start)
	builder.writeStr(" AND ")
	builder.writeSyntax(this.End)
}

func NewWindowFrameBetweenSyntax() *WindowFrameBetweenSyntax {
	s := &WindowFrameBetweenSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_WindowFrameExtentSyntax = ExtendWindowFrameExtentSyntax(s)
	return s
}
