package ast

// 函数
type I_FunctionSyntax interface {
	I_ExprSyntax
	M_FunctionSyntax_() *M_FunctionSyntax
}

type M_FunctionSyntax struct {
	I          I_FunctionSyntax
	Name       string
	Parameters *ExprListSyntax
}

func (this *M_FunctionSyntax) M_FunctionSyntax_() *M_FunctionSyntax {
	return this
}

func (this *M_FunctionSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitFunctionSyntax(this)
}

func (this *M_FunctionSyntax) AddParameter(parameter I_ExprSyntax) {
	if this.Parameters == nil {
		this.Parameters = NewExprListSyntax()
	}
	this.Parameters.Add(parameter)
}

func ExtendFunctionSyntax(i I_FunctionSyntax) *M_FunctionSyntax {
	return &M_FunctionSyntax{I: i}
}
