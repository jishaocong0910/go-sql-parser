package ast

type NullSyntax struct {
	*M_Syntax
	*M_ExprSyntax
}

func (this *NullSyntax) accept(I_Visitor) {}

func (this *NullSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("NULL")
}

func NewNullSyntax() *NullSyntax {
	s := &NullSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
