package ast

type OverSyntax struct {
	*M_Syntax
	Window I_OverWindowSyntax
}

func (this *OverSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitOverSyntax(this)
}

func (this *OverSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("OVER")
	if _, ok := this.Window.(I_IdentifierSyntax); ok {
		builder.writeSpace()
	}
	builder.writeSyntax(this.Window)
}

func NewOverSyntax() *OverSyntax {
	s := &OverSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
