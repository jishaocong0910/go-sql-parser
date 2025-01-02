package ast

type FunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
}

func (this *FunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	if this.Parameters != nil {
		builder.writeSyntaxWithFormat(this.Parameters, false)
	}
}

func NewFunctionSyntax() *FunctionSyntax {
	s := &FunctionSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	return s
}
