package ast

type NStringSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Str I_StringSyntax
}

func (this *NStringSyntax) accept(I_Visitor) {}

func (this *NStringSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("N")
	builder.writeSyntax(this.Str)
}

func NewNStringSyntax() *NStringSyntax {
	s := &NStringSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
