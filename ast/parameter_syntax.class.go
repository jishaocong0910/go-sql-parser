package ast

type ParameterSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Index int
}

func (this *ParameterSyntax) accept(I_Visitor) {}

func (this *ParameterSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("?")
}

func NewParameterSyntax() *ParameterSyntax {
	s := &ParameterSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
