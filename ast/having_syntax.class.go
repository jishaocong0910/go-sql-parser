package ast

type HavingSyntax struct {
	*M_Syntax
	Condition I_ExprSyntax
}

func (this *HavingSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitHavingSyntax(this)
}

func (this *HavingSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("HAVING")
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntax(this.Condition)
}

func NewHavingSyntax() *HavingSyntax {
	s := &HavingSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
