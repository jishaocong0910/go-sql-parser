package ast

type ExprListSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_ListSyntax[I_ExprSyntax]
	*M_ExprListSyntax[I_ExprSyntax]
}

func (this *ExprListSyntax) writeSql(builder *sqlBuilder) {
	this.M_ExprListSyntax_().writeSql(builder)
}

func (this *ExprListSyntax) IsExprList() bool {
	return this.M_ExprListSyntax_().IsExprList()
}

func (this *ExprListSyntax) ExprLen() int {
	return this.M_ExprListSyntax_().ExprLen()
}

func (this *ExprListSyntax) GetExpr(i int) I_ExprSyntax {
	return this.M_ExprListSyntax_().GetExpr(i)
}

func NewExprListSyntax() *ExprListSyntax {
	s := &ExprListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[I_ExprSyntax](s)
	s.M_ExprListSyntax = ExtendExprListSyntax[I_ExprSyntax](s)
	return s
}
