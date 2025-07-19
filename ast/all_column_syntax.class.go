package ast

type AllColumnSyntax struct {
	*Syntax__
	*SelectItemSyntax__
	*PropertyValueSyntax__
}

func (this *AllColumnSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitAllColumnSyntax(this)
}

func (this *AllColumnSyntax) Value() string {
	return "*"
}

func (this *AllColumnSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("*")
}

func NewAllColumnSyntax() *AllColumnSyntax {
	s := &AllColumnSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.SelectItemSyntax__ = ExtendSelectItemSyntax(s)
	s.PropertyValueSyntax__ = ExtendPropertyValueSyntax(s)
	return s
}
