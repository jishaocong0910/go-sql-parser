package ast

type MySqlIdentifierListSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_ListSyntax[I_IdentifierSyntax]
	*M_ExprListSyntax[I_IdentifierSyntax]
	*M_IdentifierListSyntax
}

func (this *MySqlIdentifierListSyntax) writeSql(builder *sqlBuilder) {
	this.M_ExprListSyntax_().writeSql(builder)
}

func (this *MySqlIdentifierListSyntax) IsExprList() bool {
	return this.M_ExprListSyntax_().IsExprList()
}

func (this *MySqlIdentifierListSyntax) ExprLen() int {
	return this.M_ExprListSyntax_().ExprLen()
}

func (this *MySqlIdentifierListSyntax) GetExpr(i int) I_ExprSyntax {
	return this.M_ExprListSyntax_().GetExpr(i)
}

func NewMySqlIdentifierListSyntax() *MySqlIdentifierListSyntax {
	s := &MySqlIdentifierListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[I_IdentifierSyntax](s)
	s.M_ExprListSyntax = ExtendExprListSyntax[I_IdentifierSyntax](s)
	s.M_IdentifierListSyntax = ExtendIdentifierListSyntax(s)
	return s
}
