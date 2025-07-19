package ast

// 表达式列表
type ExprListSyntax_[T ExprSyntax_] interface {
	ExprListSyntax_() *ExprListSyntax__[T]
	ExprSyntax_
	ListSyntax_[T]
}

type ExprListSyntax__[T ExprSyntax_] struct {
	I ExprListSyntax_[T]
}

func (this *ExprListSyntax__[T]) ExprListSyntax_() *ExprListSyntax__[T] {
	return this
}

func (this *ExprListSyntax__[T]) writeSql(builder *sqlBuilder) {
	this.I.Syntax_().Format = false
	this.I.ListSyntax_().writeSql(builder)
}

func (this *ExprListSyntax__[T]) IsExprList() bool {
	return true
}

func (this *ExprListSyntax__[T]) ExprLen() int {
	return this.I.ListSyntax_().Len()
}

func (this *ExprListSyntax__[T]) GetExpr(i int) ExprSyntax_ {
	return this.I.ListSyntax_().Get(i)
}

func ExtendExprListSyntax[T ExprSyntax_](i ExprListSyntax_[T]) *ExprListSyntax__[T] {
	return &ExprListSyntax__[T]{I: i}
}
