package ast

type WhereSyntax struct {
	*M_Syntax
	Condition I_ExprSyntax
}

func (this *WhereSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitWhereSyntax(this)
}

func (this *WhereSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("WHERE ")
	builder.writeSyntax(this.Condition)
}

func NewWhereSyntax() *WhereSyntax {
	s := &WhereSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
