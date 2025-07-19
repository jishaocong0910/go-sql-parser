package ast

type HavingSyntax struct {
	*Syntax__
	Condition ExprSyntax_
}

func (this *HavingSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitHavingSyntax(this)
}

func (this *HavingSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("HAVING")
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntax(this.Condition)
}

func NewHavingSyntax() *HavingSyntax {
	s := &HavingSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
