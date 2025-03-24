package ast

type SelectColumnSyntax struct {
	*M_Syntax
	*M_SelectItemSyntax
	Expr  I_ExprSyntax
	Alias I_AliasSyntax
}

func (this *SelectColumnSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitSelectColumnSyntax(this)
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_SelectItemSyntax = ExtendSelectItemSyntax(s)
	return s
}
