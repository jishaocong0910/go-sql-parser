package ast

type AggregateFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
	*M_AggregateFunctionSyntax
}

func (this *AggregateFunctionSyntax) accept(iv I_Visitor) {
	this.M_AggregateFunctionSyntax.accept(iv)
}

func NewAggregateFunctionSyntax() *AggregateFunctionSyntax {
	s := &AggregateFunctionSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.M_AggregateFunctionSyntax = ExtendAggregateFunctionSyntax(s)
	return s
}
