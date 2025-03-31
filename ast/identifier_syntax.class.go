package ast

type IdentifierSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_AliasSyntax
	*M_ColumnItemSyntax
	*M_PropertyValueSyntax
	*M_OverWindowSyntax
	*M_IdentifierSyntax
}

func (this *IdentifierSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitIdentifierSyntax(this)
}

func (this *IdentifierSyntax) Sql() string {
	return this.Name
}

func NewIdentifierSyntax() *IdentifierSyntax {
	i := &IdentifierSyntax{}
	i.M_Syntax = ExtendSyntax(i)
	i.M_ExprSyntax = ExtendExprSyntax(i)
	i.M_AliasSyntax = ExtendAliasSyntax(i)
	i.M_ColumnItemSyntax = ExtendColumnItemSyntax(i)
	i.M_PropertyValueSyntax = ExtendPropertyValueSyntax(i)
	i.M_OverWindowSyntax = ExtendOverWindowSyntax(i)
	i.M_IdentifierSyntax = ExtendIdentifierSyntax(i)
	return i
}
