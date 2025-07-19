package ast

type MySqlFalseSyntax struct {
	*Syntax__
	*ExprSyntax__
}

func (this *MySqlFalseSyntax) accept(Visitor_) {}

func (this *MySqlFalseSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("FALSE")
}

func NewMySqlFalseSyntax() *MySqlFalseSyntax {
	s := &MySqlFalseSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
