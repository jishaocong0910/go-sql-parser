package ast

type NullSyntax struct {
	*Syntax__
	*ExprSyntax__
}

func (this *NullSyntax) accept(Visitor_) {}

func (this *NullSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("NULL")
}

func NewNullSyntax() *NullSyntax {
	s := &NullSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
