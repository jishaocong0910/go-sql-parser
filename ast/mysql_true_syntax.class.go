package ast

type MySqlTrueSyntax struct {
	*M_Syntax
	*M_ExprSyntax
}

func (this *MySqlTrueSyntax) accept(I_Visitor) {}

func (this *MySqlTrueSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("TRUE")
}

func NewMySqlTrueSyntax() *MySqlTrueSyntax {
	s := &MySqlTrueSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
