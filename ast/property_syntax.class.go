package ast

type PropertySyntax struct {
	*Syntax__
	*ExprSyntax__
	*ColumnItemSyntax__
	Owner IdentifierSyntax_
	Value PropertyValueSyntax_
}

func (this *PropertySyntax) FullColumn() string {
	return this.Owner.IdentifierSyntax_().Name + "." + this.Value.Value()
}

func (this *PropertySyntax) TableAlias() string {
	return this.Owner.IdentifierSyntax_().Name
}

func (this *PropertySyntax) Column() string {
	return this.Value.Value()
}

func (this *PropertySyntax) accept(v_ Visitor_) {
	v_.visitor_().visitPropertySyntax(this)
}

func (this *PropertySyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Owner)
	builder.writeStr(".")
	builder.writeSyntax(this.Value)
}

func NewPropertySyntax() *PropertySyntax {
	s := &PropertySyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.ColumnItemSyntax__ = ExtendColumnItemSyntax(s)
	return s
}
