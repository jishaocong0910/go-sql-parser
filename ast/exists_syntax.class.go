package ast

type ExistsSyntax struct {
	*Syntax__
	*ExprSyntax__
	Query QuerySyntax_
}

func (this *ExistsSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitExistsSyntax(this)
}

func (this *ExistsSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("EXISTS")
	builder.writeSyntax(this.Query)
}

func NewExistsSyntax() *ExistsSyntax {
	s := &ExistsSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
