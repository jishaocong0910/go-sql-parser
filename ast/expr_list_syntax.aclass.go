package ast

// 表达式列表
type I_ExprListSyntax[T I_ExprSyntax] interface {
	I_ExprSyntax
	I_ListSyntax[T]
	M_5904E30AECD8() *M_ExprListSyntax[T]
}

type M_ExprListSyntax[T I_ExprSyntax] struct {
	I I_ExprListSyntax[T]
}

func (this *M_ExprListSyntax[T]) M_5904E30AECD8() *M_ExprListSyntax[T] {
	return this
}

func (this *M_ExprListSyntax[T]) writeSql(builder *sqlBuilder) {
	this.I.M_5CF6320E8474().Format = false
	this.I.M_97C5898ADC1C().writeSql(builder)
}

func (this *M_ExprListSyntax[T]) IsExprList() bool {
	return true
}

func (this *M_ExprListSyntax[T]) ExprLen() int {
	return this.I.M_97C5898ADC1C().Len()
}

func (this *M_ExprListSyntax[T]) GetExpr(i int) I_ExprSyntax {
	return this.I.M_97C5898ADC1C().Get(i)
}

func ExtendExprListSyntax[T I_ExprSyntax](i I_ExprListSyntax[T]) *M_ExprListSyntax[T] {
	return &M_ExprListSyntax[T]{I: i}
}
