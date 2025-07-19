package ast

type SelectColumnSyntax struct {
	*Syntax__
	*SelectItemSyntax__
	Expr  ExprSyntax_
	Alias AliasSyntax_
}

func (this *SelectColumnSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitSelectColumnSyntax(this)
}

func (this *SelectColumnSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Expr)
	if this.Alias != nil {
		builder.writeStr(" AS ")
		builder.writeSyntax(this.Alias)
	}
}

func NewSelectColumnSyntax() *SelectColumnSyntax {
	s := &SelectColumnSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.SelectItemSyntax__ = ExtendSelectItemSyntax(s)
	return s
}
