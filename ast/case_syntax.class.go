package ast

type CaseSyntax struct {
	*Syntax__
	*ExprSyntax__
	ValueExpr    ExprSyntax_
	WhenItemList *CaseWhenItemListSyntax
	ElseExr      ExprSyntax_
}

func (this *CaseSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitCaseSyntax(this)
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
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
