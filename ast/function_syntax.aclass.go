package ast

// 函数
type FunctionSyntax_ interface {
	FunctionSyntax_() *FunctionSyntax__
	ExprSyntax_
}

type FunctionSyntax__ struct {
	I          FunctionSyntax_
	Name       string
	Parameters *ExprListSyntax
}

func (this *FunctionSyntax__) FunctionSyntax_() *FunctionSyntax__ {
	return this
}

func (this *FunctionSyntax__) accept(v_ Visitor_) {
	v_.visitor_().visitFunctionSyntax__(this)
}

func (this *FunctionSyntax__) AddParameter(parameter ExprSyntax_) {
	if this.Parameters == nil {
		this.Parameters = NewExprListSyntax()
	}
	this.Parameters.Add(parameter)
}

func ExtendFunctionSyntax(i FunctionSyntax_) *FunctionSyntax__ {
	return &FunctionSyntax__{I: i}
}
