package ast

type CaseSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	ValueExpr    I_ExprSyntax
	WhenItemList *CaseWhenItemListSyntax
	ElseExr      I_ExprSyntax
}

func (this *CaseSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitCaseSyntax(this)
}

func (this *CaseSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("CASE")
	if this.ValueExpr != nil {
		builder.writeSpace()
		builder.writeSyntax(this.ValueExpr)
	}
	builder.writeSpaceOrLf(this, false)
	builder.writeSyntax(this.WhenItemList)
	if this.ElseExr != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("ELSE ")
		builder.writeSyntax(this.ElseExr)
	}
	builder.writeSpace()
	builder.writeStr("END")
}

func NewCaseSyntax() *CaseSyntax {
	s := &CaseSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
