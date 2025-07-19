package ast

type OverSyntax struct {
	*Syntax__
	Window OverWindowSyntax_
}

func (this *OverSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitOverSyntax(this)
}

func (this *OverSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("OVER")
	if _, ok := this.Window.(IdentifierSyntax_); ok {
		builder.writeSpace()
	}
	builder.writeSyntax(this.Window)
}

func NewOverSyntax() *OverSyntax {
	s := &OverSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
