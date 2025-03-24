package ast

type PropertySyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_ColumnItemSyntax
	Owner I_IdentifierSyntax
	Value I_PropertyValueSyntax
}

func (this *PropertySyntax) FullColumn() string {
	return this.Owner.M_IdentifierSyntax_().Name + "." + this.Value.Value()
}

func (this *PropertySyntax) TableAlias() string {
	return this.Owner.M_IdentifierSyntax_().Name
}

func (this *PropertySyntax) Column() string {
	return this.Value.Value()
}

func (this *PropertySyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitPropertySyntax(this)
}

func (this *PropertySyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Owner)
	builder.writeStr(".")
	builder.writeSyntax(this.Value)
}

func NewPropertySyntax() *PropertySyntax {
	s := &PropertySyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_ColumnItemSyntax = ExtendColumnItemSyntax(s)
	return s
}
