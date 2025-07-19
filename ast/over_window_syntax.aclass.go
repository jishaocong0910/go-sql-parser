package ast

// 窗口函数中的OVER语法
type OverWindowSyntax_ interface {
	OverWindowSyntax_() *OverWindowSyntax__
	Syntax_
}

type OverWindowSyntax__ struct {
	I OverWindowSyntax_
}

func (this *OverWindowSyntax__) OverWindowSyntax_() *OverWindowSyntax__ {
	return this
}

func ExtendOverWindowSyntax(i OverWindowSyntax_) *OverWindowSyntax__ {
	return &OverWindowSyntax__{I: i}
}
