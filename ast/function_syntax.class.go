package ast

type FunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
}

func (this *FunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	if this.Parameters != nil {
		builder.writeSyntaxWithFormat(this.Parameters, false)
	}
}

func NewFunctionSyntax() *FunctionSyntax {
	s := &FunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	return s
}
