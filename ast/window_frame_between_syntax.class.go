package ast

type WindowFrameBetweenSyntax struct {
	*Syntax__
	*WindowFrameExtentSyntax__
	Start WindowFrameStartEndSyntax_
	End   WindowFrameStartEndSyntax_
}

func (this *WindowFrameBetweenSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitWindowFrameBetweenSyntax(this)
}

func (this *WindowFrameBetweenSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("BETWEEN ")
	builder.writeSyntax(this.Start)
	builder.writeStr(" AND ")
	builder.writeSyntax(this.End)
}

func NewWindowFrameBetweenSyntax() *WindowFrameBetweenSyntax {
	s := &WindowFrameBetweenSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.WindowFrameExtentSyntax__ = ExtendWindowFrameExtentSyntax(s)
	return s
}
