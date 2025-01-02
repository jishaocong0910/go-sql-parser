package ast

type MySqlFalseSyntax struct {
	*M_Syntax
	*M_ExprSyntax
}

func (this *MySqlFalseSyntax) accept(I_Visitor) {}

func (this *MySqlFalseSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("FALSE")
}

func NewMySqlFalseSyntax() *MySqlFalseSyntax {
	s := &MySqlFalseSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
