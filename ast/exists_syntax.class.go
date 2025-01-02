package ast

type ExistsSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Query I_QuerySyntax
}

func (this *ExistsSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitExistsSyntax(this)
}

func (this *ExistsSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("EXISTS")
	builder.writeSyntax(this.Query)
}

func NewExistsSyntax() *ExistsSyntax {
	s := &ExistsSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
