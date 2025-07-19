package ast

type ExprListSyntax struct {
	*Syntax__
	*ExprSyntax__
	*ListSyntax__[ExprSyntax_]
	*ExprListSyntax__[ExprSyntax_]
}

func (this *ExprListSyntax) writeSql(builder *sqlBuilder) {
	this.ExprListSyntax_().writeSql(builder)
}

func (this *ExprListSyntax) IsExprList() bool {
	return this.ExprListSyntax_().IsExprList()
}

func (this *ExprListSyntax) ExprLen() int {
	return this.ExprListSyntax_().ExprLen()
}

func (this *ExprListSyntax) GetExpr(i int) ExprSyntax_ {
	return this.ExprListSyntax_().GetExpr(i)
}

func NewExprListSyntax() *ExprListSyntax {
	s := &ExprListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[ExprSyntax_](s)
	s.ExprListSyntax__ = ExtendExprListSyntax[ExprSyntax_](s)
	return s
}
