package ast

type CaseWhenItemSyntax struct {
	*Syntax__
	Condition ExprSyntax_
	Result    ExprSyntax_
}

func (this *CaseWhenItemSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitCaseWhenItemSyntax(this)
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
	s.Syntax__ = ExtendSyntax(s)
	return s
}
