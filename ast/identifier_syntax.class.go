package ast

type IdentifierSyntax struct {
	*Syntax__
	*ExprSyntax__
	*AliasSyntax__
	*ColumnItemSyntax__
	*PropertyValueSyntax__
	*OverWindowSyntax__
	*IdentifierSyntax__
}

func (this *IdentifierSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitIdentifierSyntax(this)
}

func (this *IdentifierSyntax) Sql() string {
	return this.Name
}

func NewIdentifierSyntax() *IdentifierSyntax {
	i := &IdentifierSyntax{}
	i.Syntax__ = ExtendSyntax(i)
	i.ExprSyntax__ = ExtendExprSyntax(i)
	i.AliasSyntax__ = ExtendAliasSyntax(i)
	i.ColumnItemSyntax__ = ExtendColumnItemSyntax(i)
	i.PropertyValueSyntax__ = ExtendPropertyValueSyntax(i)
	i.OverWindowSyntax__ = ExtendOverWindowSyntax(i)
	i.IdentifierSyntax__ = ExtendIdentifierSyntax(i)
	return i
}
