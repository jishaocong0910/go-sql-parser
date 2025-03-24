package ast

type OrderBySyntax struct {
	*M_Syntax
	OrderByItemList *OrderingItemListSyntax
}

func (this *OrderBySyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitOrderBySyntax(this)
}

func (this *OrderBySyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("ORDER BY")
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntaxWithFormat(this.OrderByItemList, this.Format)
}

func NewOrderBySyntax() *OrderBySyntax {
	s := &OrderBySyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
