package ast

type CaseWhenItemSyntax struct {
	*M_Syntax
	Condition I_ExprSyntax
	Result    I_ExprSyntax
}

func (this *CaseWhenItemSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitCaseWhenItemSyntax(this)
}

func (this *CaseWhenItemSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("WHEN ")
	builder.writeSyntax(this.Condition)
	builder.writeSpace()
	builder.writeStr("THEN ")
	builder.writeSyntax(this.Result)
}

func NewCaseWhenItem() *CaseWhenItemSyntax {
	s := &CaseWhenItemSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
