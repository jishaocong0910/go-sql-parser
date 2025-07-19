package ast

type NStringSyntax struct {
	*Syntax__
	*ExprSyntax__
	Str StringSyntax_
}

func (this *NStringSyntax) accept(Visitor_) {}

func (this *NStringSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("N")
	builder.writeSyntax(this.Str)
}

func NewNStringSyntax() *NStringSyntax {
	s := &NStringSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
