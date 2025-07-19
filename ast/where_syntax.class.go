package ast

type WhereSyntax struct {
	*Syntax__
	Condition ExprSyntax_
}

func (this *WhereSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitWhereSyntax(this)
}

func (this *WhereSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("WHERE ")
	builder.writeSyntax(this.Condition)
}

func NewWhereSyntax() *WhereSyntax {
	s := &WhereSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
