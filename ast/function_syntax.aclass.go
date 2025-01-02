package ast

// 函数
type I_FunctionSyntax interface {
	I_ExprSyntax
	M_9070BBA0A777() *M_FunctionSyntax
}

type M_FunctionSyntax struct {
	I          I_FunctionSyntax
	Name       string
	Parameters *ExprListSyntax
}

func (this *M_FunctionSyntax) M_9070BBA0A777() *M_FunctionSyntax {
	return this
}

func (this *M_FunctionSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitFunctionSyntax(this)
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
