package ast

type MySqlIdentifierListSyntax struct {
	*Syntax__
	*ExprSyntax__
	*ListSyntax__[IdentifierSyntax_]
	*ExprListSyntax__[IdentifierSyntax_]
	*IdentifierListSyntax__
}

func (this *MySqlIdentifierListSyntax) writeSql(builder *sqlBuilder) {
	this.ExprListSyntax_().writeSql(builder)
}

func (this *MySqlIdentifierListSyntax) IsExprList() bool {
	return this.ExprListSyntax_().IsExprList()
}

func (this *MySqlIdentifierListSyntax) ExprLen() int {
	return this.ExprListSyntax_().ExprLen()
}

func (this *MySqlIdentifierListSyntax) GetExpr(i int) ExprSyntax_ {
	return this.ExprListSyntax_().GetExpr(i)
}

func NewMySqlIdentifierListSyntax() *MySqlIdentifierListSyntax {
	s := &MySqlIdentifierListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[IdentifierSyntax_](s)
	s.ExprListSyntax__ = ExtendExprListSyntax[IdentifierSyntax_](s)
	s.IdentifierListSyntax__ = ExtendIdentifierListSyntax(s)
	return s
}
