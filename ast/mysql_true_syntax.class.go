package ast

type MySqlTrueSyntax struct {
	*Syntax__
	*ExprSyntax__
}

func (this *MySqlTrueSyntax) accept(Visitor_) {}

func (this *MySqlTrueSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("TRUE")
}

func NewMySqlTrueSyntax() *MySqlTrueSyntax {
	s := &MySqlTrueSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
