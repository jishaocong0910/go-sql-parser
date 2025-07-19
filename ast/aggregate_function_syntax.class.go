package ast

type AggregateFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	*AggregateFunctionSyntax__
}

func (this *AggregateFunctionSyntax) accept(v_ Visitor_) {
	this.AggregateFunctionSyntax__.accept(v_)
}

func NewAggregateFunctionSyntax() *AggregateFunctionSyntax {
	s := &AggregateFunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.AggregateFunctionSyntax__ = ExtendAggregateFunctionSyntax(s)
	return s
}
