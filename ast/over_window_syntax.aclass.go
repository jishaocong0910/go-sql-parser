package ast

// 窗口函数中的OVER语法
type I_OverWindowSyntax interface {
	I_Syntax
	M_OverWindowSyntax_() *M_OverWindowSyntax
}

type M_OverWindowSyntax struct {
	I I_OverWindowSyntax
}

func (this *M_OverWindowSyntax) M_OverWindowSyntax_() *M_OverWindowSyntax {
	return this
}

func ExtendOverWindowSyntax(i I_OverWindowSyntax) *M_OverWindowSyntax {
	return &M_OverWindowSyntax{I: i}
}
